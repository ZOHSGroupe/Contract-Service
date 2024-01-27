pipeline {
    agent any
    tools {
        maven 'Maven'
    }
    environment {
        // Customize these variables based on your project
        DOCKER_HUB_REPO = 'ouail02/contract-service'
        SPRING_PROFILES_ACTIVE = 'production'
        BRANCH = 'main'
        Git_REPO = 'https://github.com/ZOHSGroupe/Contract-Service'
    }

    stages {
        stage('Checkout') {
            steps {
                // Manually specify the Git repository URL and branch
                script {
                    checkout([$class: 'GitSCM', branches: [[name: "${BRANCH}"]], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: "${Git_REPO}"]]])
                }
            }
        }
        stage('Build Docker Image') {
            steps {
                sh "docker build -t ${DOCKER_HUB_REPO}:${env.BUILD_NUMBER} ."
            }
        }
        stage('Push to Docker Hub') {
            steps {
                script {
                    // Log in to Docker Hub
                    withCredentials([string(credentialsId: 'DOCKER_HUB_USERNAME', variable: 'DOCKER_HUB_USERNAME'), string(credentialsId: 'DOCKER_HUB_PASSWORD', variable: 'DOCKER_HUB_PASSWORD')]) {
                        sh "docker login -u ${DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_PASSWORD}"
                    }
                    sh "docker push ${DOCKER_HUB_REPO}:${env.BUILD_NUMBER}"
                }
            }
        }

        // Add your other stages here

        // stage('Deploy to k8s') {
        //     steps {
        //         script {
        //             kubernetesDeploy (configs: 'deploymentservice.yaml',kubeconfigId: 'k8sconfigpwd')
        //         }
        //     }
        // }

    }
    post {
            always {
                // Clean up: Remove the Docker image locally
                script {
                    try {
                        // Your existing post-build actions here
                    } finally {
                        // Cleanup: Remove the Docker image locally
                        sh "docker rmi ${DOCKER_HUB_REPO}:${env.BUILD_NUMBER}"
                    }
                }
            }
        }
}
