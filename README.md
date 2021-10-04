# nus-hello-word

In this lab, we are building a continuous delivery (CD) pipeline. We are using a very simple application written in Go. For the sake of simplicity, we are going to run only one type of test against the code. The prerequisites for this lab are as follows:

1) A running Jenkins instance. 
2) Image registry like ECR or GCR
3) An account on GitHub

Step 1: The Application Files

main.go
Our sample application will respond with ‘Hello World’ to any GET request. 

main_test.go
Since we are building a CD pipeline, we should have some tests in place. Our code is so simple that it only needs one test case; ensuring that we receive the correct string when we hit the root URL.

Dockerfile
This is where we package our application

service.yml
we need at least a service and a deployment since we are using Kubernetes as the platform on which we host this application.

deployment.yml
The application itself, once dockerized, can be deployed to Kubernetes through a Deployment resource. The deployment.yml file looks as follows:

Step 2:
a) Create a Kubernetes Cluster in GCP
b) Setup Helm
c) Configure and install Jenkins
d) Test is pod running

Step 3: Create Jenkins Pipeline Job
1) We used the Poll SCM as the build trigger; setting this option instructs Jenkins to check the Git repository on a periodic basis (every minute as indicated by * * * * *). If the repo has changed since the last poll, the job is triggered.
2) In the pipeline itself, we specified the repository URL and the credentials. The branch is main.
3) We are adding all the job’s code in a Jenkinsfile that is stored in the same repository as the code.

Step 4: Configure Jenkins Credentials For GitHub and Docker Hub

Step 5: Create the JenkinsFile
Basically, the pipeline contains four stages:
1) Build is where we build the Go binary and ensure that nothing is wrong in the build process.
2) Test is where we apply a simple UAT test to ensure that the application works as expected.
3) Publish, where the Docker image is built and pushed to the registry. After that, any environment can make use of it.
4) Deploy, this is the final step where Ansible is invoked to contact Kubernetes and apply the definition files.

Step 6: Testing CD Pipeline
1) Commit our code to GitHub and ensure that our code moves through the pipeline until it reaches the cluster:
a) commit and push the code to GitHub
b) On Jenkins, we can either wait for the job to get triggered automatically, or we can just click on “Build Now”.
c) When the job succeed, get the node IP address and initiate an HTTP request to our app and should show "hello word" message

Step 7: Test Error
1) Go to main.go and change the message to "Hello Word!!!!!" and push the change to Github.
2) Pipeline should stop at the Test stage and this proves that pipeline will not ship faulty code to the target environment.

Summary
1) Through Jenkins, we were able to pull the code from the repository, build and test it using a relevant Docker Image.
2) Next, we Dockerize and push our application - since it passed our tests - to Docker Hub.
3) Using Jenkins pipelines and Ansible makes it very easy and flexible to change the workflow with very little friction. For example, we can add more tests to the Test stage, we can change the version of Go that’s used to build and test the code, and we can use more variables to change other aspects in deployment and service definitions.
4) The best part here is that we are using Kubernetes deployments, which ensures that we have zero downtime for the application when we are changing the container image. This is possible because Deployments use the rolling update method by default to terminate and recreate containers one at a time. Only when the new container is up and healthy does the Deployment terminate the old one.



