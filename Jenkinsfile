pipeline {
    agent { label 'ubuntu' }
    tools {
      go '1.19.6'
      nodejs '18.10.0'
    }
        stages {
        stage('Cloning git reposiry') {
            steps {
                git branch: 'release/v1.17', url: 'git@github.com:bortnychuk/gitea.git'
            }
        }
        stage('Configuring Go, nodejs and build-essential') {
            steps {
                sh 'npm version'
                sh 'go version'
                sh 'sudo apt install make'
                sh 'sudo apt-get install build-essential'
                sh 'sudo go install github.com/google/go-licenses@latest'
                sh 'ls -la'
            }
        }
        stage('Formating code') {
            steps {
                sh 'make fmt'
            }
        }
        stage('Building Gitea application') {
            steps {
                sh 'TAGS="bindata sqlite sqlite_unlock_notify" make build'
            }
        }
        stage('Running tests') {
            steps {
                sh 'TAGS="bindata sqlite sqlite_unlock_notify" make test'
            }
        }
        stage('Publishing and runinig Gitea on Test Server') {
            steps {
                sshPublisher(publishers: [sshPublisherDesc(configName: 'MyGiteaServer', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: 'sudo systemctl start gitea', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '/usr/gitEA', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '*')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
            }
        }
        stage('Checking connection with Gitea test') {
            steps {
                sh 'sleep 10'
                sh 'curl -I --http2 -s 192.168.1.34:3000 | grep HTTP'
            }
        }
    }        
}
