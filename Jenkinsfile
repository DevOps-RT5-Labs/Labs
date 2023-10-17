pipeline {
  agent any
  

  environment {
    DOCKERHUB_CREDENTIALS = credentials('dockerhub')
  }

  stages {
      stage('Build Docker Image') {
        when {
            branch "dev"
        }

          steps {
            echo 'building ...'
            sh 'docker build -t devops-demo-app:latest 1-docker/apps'
         }
      }



      stage('deploy docker images!') {
           when {
              branch "test"
          }
          steps {
            echo 'Login to DockerHub'
            sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'

            echo 'building ...'
            sh 'docker build -t devops-demo-app:latest .'
         }
      }
  }

    post {
        always {
            sh 'docker logout'
        }
    }
}