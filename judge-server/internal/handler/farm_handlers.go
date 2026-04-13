package handler

import (
	"encoding/json"
	"fmt"
	"judge-server/internal/model"
	"judge-server/internal/service"
	"net/http"
	"strings"
)

// FarmHTTPHandler 处理农业相关的 HTTP 请求
type FarmHTTPHandler struct {
	farmService *service.FarmService
}

// NewFarmHTTPHandler 创建新的农业处理器
func NewFarmHTTPHandler(farmService *service.FarmService) *FarmHTTPHandler {
	return &FarmHTTPHandler{
		farmService: farmService,
	}
}

// HandleFarm 处理 /api/farm 相关端点
func (h *FarmHTTPHandler) HandleFarm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 提取基地ID和用户ID
	baseID := r.URL.Query().Get("base_id")
	githubIDStr := r.URL.Query().Get("github_id")

	if baseID == "" || githubIDStr == "" {
		http.Error(w, "base_id and github_id parameters required", http.StatusBadRequest)
		return
	}

	var githubID int
	_, err := fmt.Sscanf(githubIDStr, "%d", &githubID)
	if err != nil {
		http.Error(w, "invalid github_id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getFarmStatus(w, baseID)
	case http.MethodPost:
		h.handleFarmAction(w, r, baseID, githubID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getFarmStatus 获取菜园状态
func (h *FarmHTTPHandler) getFarmStatus(w http.ResponseWriter, baseID string) {
	plots, err := h.farmService.GetFarmStatus(baseID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get farm status: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"plots":    plots,
		"capacity": 8,
	})
}

// handleFarmAction 处理农业行为（种植/收获）
func (h *FarmHTTPHandler) handleFarmAction(w http.ResponseWriter, r *http.Request, baseID string, githubID int) {
	action := r.URL.Query().Get("action")

	switch action {
	case "plant":
		h.plantCrop(w, r, baseID, githubID)
	case "harvest":
		h.harvestCrop(w, r, baseID)
	case "eat":
		h.eatCrop(w, r, baseID)
	default:
		http.Error(w, "unknown action. Use 'plant', 'harvest', or 'eat'", http.StatusBadRequest)
	}
}

// plantCrop 种植作物
func (h *FarmHTTPHandler) plantCrop(w http.ResponseWriter, r *http.Request, baseID string, githubID int) {
	var req model.PlantCropRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}

	// 验证地块号
	if req.PlotNumber < 1 || req.PlotNumber > 8 {
		http.Error(w, "plot_number must be between 1 and 8", http.StatusBadRequest)
		return
	}

	// 这里需要获取用户等级，暂时假设为 5
	userLevel := 5

	plot, err := h.farmService.PlantCrop(baseID, githubID, req.PlotNumber, req.CropType, req.Quantity, userLevel)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to plant crop: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"plot":    plot,
	})
}

// harvestCrop 收获作物
func (h *FarmHTTPHandler) harvestCrop(w http.ResponseWriter, r *http.Request, baseID string) {
	var req model.HarvestCropRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}

	if req.PlotNumber < 1 || req.PlotNumber > 8 {
		http.Error(w, "plot_number must be between 1 and 8", http.StatusBadRequest)
		return
	}

	harvest, err := h.farmService.HarvestCrop(baseID, req.PlotNumber)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to harvest crop: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"harvest": harvest,
		"message": "Crop harvested successfully",
	})
}

// eatCrop 进食作物
func (h *FarmHTTPHandler) eatCrop(w http.ResponseWriter, r *http.Request, baseID string) {
	var req model.EatFarmRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}

	eating, err := h.farmService.AllowEating(baseID, req.VisitorGitHubID, req.CropType, req.VisitorPetID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to eat: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"eating":  eating,
	})
}

// HandleBaseVisit 处理基地访问
func (h *FarmHTTPHandler) HandleBaseVisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从路径提取 base_id: /api/base/{base_id}/visit
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/base/"), "/")
	if len(pathParts) < 2 {
		http.Error(w, "base_id required", http.StatusBadRequest)
		return
	}
	baseID := pathParts[0]

	// 从查询参数获取访问者 ID
	visitorIDStr := r.URL.Query().Get("visitor_id")
	if visitorIDStr == "" {
		http.Error(w, "visitor_id parameter required", http.StatusBadRequest)
		return
	}

	var visitorID int
	_, err := fmt.Sscanf(visitorIDStr, "%d", &visitorID)
	if err != nil {
		http.Error(w, "invalid visitor_id", http.StatusBadRequest)
		return
	}

	var req model.VisitBaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}

	// TODO: Get owner github_id from baseID lookup
	// For now, we'll assume it's provided or we look it up from the farm
	// Get the base owner from farm data
	ownerGitHubID, err := h.getBaseOwner(baseID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get base owner: %v", err), http.StatusInternalServerError)
		return
	}

	// For now, simulate a battle result (50% win/loss)
	// In production, this would call the actual battle engine
	battleResult := model.BattleResultWin
	if visitorID%2 == 0 {
		battleResult = model.BattleResultLose
	}

	// Calculate rewards based on visit type and battle result
	visitorPoints, visitorReputation, ownerPoints, ownerReputation, _ := service.CalculateRewards(req.VisitType, battleResult)

	// Record the visit
	visit, err := h.farmService.RecordBaseVisit(baseID, ownerGitHubID, visitorID, req.VisitType, req.VisitorPetID, "", battleResult, visitorPoints, visitorReputation, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to record visit: %v", err), http.StatusInternalServerError)
		return
	}

	// Update both players' reputation
	if err := h.farmService.UpdateReputation(visitorID, visitorPoints, visitorReputation); err != nil {
		fmt.Printf("warning: failed to update visitor reputation: %v\n", err)
	}
	if err := h.farmService.UpdateReputation(ownerGitHubID, ownerPoints, ownerReputation); err != nil {
		fmt.Printf("warning: failed to update owner reputation: %v\n", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Visit recorded",
		"visit":   visit,
	})
}

// getBaseOwner 从基地ID获取所有者的github_id
func (h *FarmHTTPHandler) getBaseOwner(baseID string) (int, error) {
	// 从base_farm表中查询
	plots, err := h.farmService.GetFarmStatus(baseID)
	if err != nil {
		return 0, fmt.Errorf("failed to get farm status: %w", err)
	}

	if len(plots) > 0 {
		return plots[0].GitHubID, nil
	}

	// If no plots, we can't determine owner
	// This is a limitation - we should have a separate bases table
	return 0, fmt.Errorf("no plots found for base, cannot determine owner")
}

// HandleReputation 处理声望查询
func (h *FarmHTTPHandler) HandleReputation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	githubIDStr := r.URL.Query().Get("github_id")
	if githubIDStr == "" {
		http.Error(w, "github_id parameter required", http.StatusBadRequest)
		return
	}

	var githubID int
	_, err := fmt.Sscanf(githubIDStr, "%d", &githubID)
	if err != nil {
		http.Error(w, "invalid github_id", http.StatusBadRequest)
		return
	}

	rep, err := h.farmService.GetReputation(githubID)
	if err != nil {
		// Try to initialize if not found
		rep, err = h.farmService.InitializeReputation(githubID)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get or initialize reputation: %v", err), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":          true,
		"reputation":       rep,
		"reputation_level": getReputationLevel(rep.TotalReputation),
	})
}

// getReputationLevel 获取声望等级描述
func getReputationLevel(reputation int) string {
	if reputation >= 100 {
		return "Legend"
	} else if reputation >= 50 {
		return "Respected"
	} else if reputation >= 20 {
		return "Known"
	} else if reputation >= 0 {
		return "Novice"
	} else {
		return "Villain"
	}
}
