pipeline {
    agent any
    stages {
        stage('Build - Sandbox') {
            when{
                expression {
                    return env.GIT_BRANCH == "origin/master"
                }
            }
            environment {
                ENV='dev'
            }
            steps {
                echo 'Building Sandbox'
                sh 'ENV=dev docker-compose up --force-recreate'
            }
        }
        stage('Test - Sandbox') {
            steps {
                echo 'Runing Test on sandbox'
            }
        }
        stage('Deploy - Sandbox') {
            when{
                expression {
                    return env.GIT_BRANCH == "origin/master"
                }
            }
            environment {
                ENV='dev'
            }
            steps {
                dir('bin'){
                    echo 'Deploying ...'
                    sh 'DIRPATH=${pwd()} docker-compose up --force-recreate -d'
                }
            }
        }
    }
}
