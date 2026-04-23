pipeline {
    agent any

    environment {
        IMAGE_NAME = 'mi-web-uoc'
        IMAGE_TAG = 'latest'
        KUBECONFIG = '/var/lib/jenkins/.kube/config'
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/adimaggiopa-ui/mi-web-go.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .'
                sh 'docker save ${IMAGE_NAME}:${IMAGE_TAG} -o ${IMAGE_NAME}.tar'
                sh 'sudo ctr -n k8s.io images import ${IMAGE_NAME}.tar'
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl apply -f k8s/deployment.yaml'
                sh 'kubectl apply -f k8s/service.yaml'
                sh 'kubectl rollout restart deployment/mi-web-deployment'
            }
        }
    }

    post {
        success {
            echo 'Despliegue completado con exito!'
        }
        failure {
            echo 'El pipeline ha fallado. Revisa los logs.'
        }
    }
}