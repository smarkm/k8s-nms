deploy=nginx-deploy.yaml
service=nginx-service.yaml

gen-nginx:
	kubectl create deployment nginx --image=nginx:1.17  --dry-run -o yaml > ${deploy}
gen-nginx-service: # please noted the label match the pods
	kubectl create service clusterip --tcp 80:80 nginx --dry-run -o yaml > ${service}

deploy-nginx-service: gen-nginx gen-nginx-service
	kubectl apply -f ${deploy}
	kubectl apply -f ${service}
	rm ${deploy} ${service}
undeloy:
	kubectl delete service nginx
	kubectl delete deploy nginx