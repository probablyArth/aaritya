package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Me(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	user, err := u.ctx.UserRepository.GetUserByID(userId.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": "Failed to fetch user data"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *UserHandler) Quizes(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	limit := 10
	quizzes, err := u.ctx.QuizRepository.GetRecentQuizzesByUserID(userId.(uint), limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": "Failed to fetch quizzes"})
		return
	}
	ctx.JSON(http.StatusOK, quizzes)
}

func (u *UserHandler) Leaderboard(ctx *gin.Context) {
	limit := 10
	users, err := u.ctx.UserRepository.GetTopUsers(limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": "Failed to fetch leaderboard"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *UserHandler) Rank(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	rank, err := u.ctx.UserRepository.GetUserRank(userId.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": "Failed to fetch user rank"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {"rank": rank})
}

func (u *UserHandler) Total(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	score, err := u.ctx.UserRepository.GetUserAverageScore(userId.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": "Failed to fetch user average score"})
		return
	}
	ctx.JSON(http.StatusOK, score)
}

func (u *UserHandler) Average(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	total, err := u.ctx.UserRepository.GetUserTotalQuizzes(userId.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": "Failed to fetch user total quizzes"})
		return
	}
	ctx.JSON(http.StatusOK, total)
}
