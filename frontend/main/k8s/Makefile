.PHONY: apply-frontend-deployment
apply-frontend-deployment:
	kubectl apply -f frontend-deployment.yaml -n app-ns

.PHONY: apply-frontend-service
apply-frontend-service:
	kubectl apply -f frontend-service.yaml -n app-ns

.PHONY: stern-frontend
stern-frontend:
	stern frontend -n app-ns
