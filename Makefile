
swagger: 
	swag init -g backend/internal/router/router.go -o backend/internal/router/doc

.PHONY:backend
backend:
	cd backend && go run cmd/root.go

.PHONY:frontend
frontend:
	cd frontend && npm run dev

