# nus-hello-word

In this lab, we are building a continuous integration/continouus delivery (CU/CD) pipeline. We are using a very simple application written in Go. For the sake of simplicity, we are going to run only one type of test against the code. The prerequisites for this setup are as follows:

1. An account on ***Github***
2. An account on ***DockerHub***
3. A running ***Jenkins*** instance
4. A running ***Kubernetes*** Cluster

## Step 1: The Application Files
#### main.go
Our sample application will respond with ‘Hello World’ to any GET request. 

### main_test.go
Since we are building a CI/CD pipeline, we should have some tests in place to ensure that we receive the correct text when we hit the root URL.

### Dockerfile
This is where we package our application.

### service.yml
we need at least a service and a deployment since we are using Kubernetes as the platform on which we host this application.

### deployment.yml
The application itself, once dockerized, can be deployed to Kubernetes through a Deployment resource. The deployment.yml file looks as follows:

## Step 2: The GitHub Repository
1. Create a repository to host and versioing our source code.

## Step 3: Kubernetes Cluster
1. Create a Google Kubernetes Engine Cluster 

## Step 4: The Jenkins Instance
1. Install Openshift Python
2. Install Ansible
3. Install Jenkins
4. Install Docker

## Step 5: The Jenkins Pipeline
1. We specified the repository URL and the credentials and point to "main" branch
2. We used the Poll SCM as the build trigger to instructs Jenkins to check the Git repository on a periodic basis
2. We are adding all the job’s code in a Jenkinsfile that is stored in the same repository as the code.

### Basically, the pipeline contains four stages:
1. Build is where we build the Go binary and ensure that nothing is wrong in the build process.
2. Test is where we apply a simple UAT test to ensure that the application works as expected.
3. Publish, where the Docker image is built and pushed to the registry. After that, any environment can make use of it.
4. Deploy, this is the final step where Ansible is invoked to contact Kubernetes and apply the definition files.

## Step 6: Configure Jenkins Credentials For GitHub, Docker Hub and Kubernetes Cluster
This setup is to allow Jenkins to connect to GitHub, Docker Hub and Kubernetes Cluster

## Step 7: Testing CI/CD Pipeline
1. Commit our code to GitHub and ensure that our code moves through the pipeline until it reaches the cluster.
3. On Jenkins, we can either wait for the job to get triggered automatically, or we can just click on “Build Now”.
4. When the job succeed, get the node IP address and initiate an HTTP request to our app and should show "hello word" message

## Step 8: Test For Error
1. Go to main.go and change the message to "Hello Word!!!!!" and push the change to Github.
2. Pipeline should stop at the Test stage and this proves that pipeline will not ship faulty code to the target environment.

## Summary
- Through Jenkins, we were able to pull the code from the repository, build and test it using a relevant Docker Image.
- Next, we Dockerize and push our application - since it passed our tests - to Docker Hub.
- Using Jenkins pipelines and Ansible makes it very easy and flexible to change the workflow with very little friction. For example, we can add more tests to the Test stage, we can change the version of Go that’s used to build and test the code, and we can use more variables to change other aspects in deployment and service definitions.
- The best part here is that we are using Kubernetes deployments, which ensures that we have zero downtime for the application when we are changing the container image. This is possible because Deployments use the rolling update method by default to terminate and recreate containers one at a time. Only when the new container is up and healthy does the Deployment terminate the old one.



