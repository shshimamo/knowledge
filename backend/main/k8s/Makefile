.PHONY: apply-migration-job
apply-migration-job:
	kubectl apply -f migration-job.yaml -n app-ns

.PHONY: apply-backend-deployment
apply-backend-deployment:
	kubectl apply -f backend-deployment.yaml -n app-ns

.PHONY: apply-backend-service
apply-backend-service:
	kubectl apply -f backend-service.yaml -n app-ns

.PHONY: stern-backend
stern-backend:
	stern backend -n app-ns
