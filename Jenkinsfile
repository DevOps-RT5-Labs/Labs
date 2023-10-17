pipeline {
  agent any
  

  environment {
    DOCKERHUB_CREDENTIALS = credentials('dockerhub')
    SNYK_TOKEN = credentials('snyk')
  }

  stages {
      stage('Build Docker Image') {
        when {
            branch "dev"
        }

          steps {
            echo 'building ...'
            sh 'HASH=$(git rev-parse --short HEAD) && docker build -t niemandx/devops-demo-app:$HASH 1-docker/apps'
         }
      }

      stage('scan for vulnerable packages!') {
          when {
              branch "dev"
          }
          steps {
            echo 'Login to Snyk'
            sh 'snyk auth $SNYK_TOKEN'

            echo 'scanning ...'
            sh 'snyk test'
         }
      }


      stage('deploy docker images!') {
           when {
              branch "dev"
          }
          steps {
            echo 'Login to DockerHub'
            sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'

            echo 'building ...'
            sh 'HASH=$(git rev-parse --short HEAD) && docker push niemandx/devops-demo-app:$HASH'
         }
      }
  }

    post {
        always {
            sh 'docker logout'
        }
    }
}