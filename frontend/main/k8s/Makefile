.PHONY: apply-deployment
apply-deployment:
	kubectl apply -f deployment.yaml -n app-ns

.PHONY: apply-service
apply-service:
	kubectl apply -f service.yaml -n app-ns

.PHONY: stern-frontend
stern-frontend:
	stern frontend -n app-ns
