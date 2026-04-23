pipeline {
    agent any

    environment {
        IMAGE_NAME = 'mi-web-uoc'
        IMAGE_TAG = 'latest'
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
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl apply -f k8s/deployment.yaml'
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
