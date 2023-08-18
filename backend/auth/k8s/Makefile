.PHONY: apply-migration-job
apply-migration-job:
	kubectl apply -f migration-job.yaml -n app-ns

.PHONY: apply-deployment
apply-deployment:
	kubectl apply -f deployment.yaml -n app-ns

.PHONY: apply-service
apply-service:
	kubectl apply -f service.yaml -n app-ns

.PHONY: stern-backend
stern-backend:
	stern backend -n app-ns
