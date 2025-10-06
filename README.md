# NBIS Assignment

Assignment for the position as backend system developer with focus on security
with reference number: UFV-PA 2025/2222.

A pull request (PR) called "Feedback" will automatically be created on your repository. You can use that PRs for questions and to request feedback. Please tag the team @NBISweden/sysdev-recruitment in the text.

## How to submit

All tasks should be done in the form of PRs towards the repo. It is up to you to decide how you will split your answers into different PRs. You can assume that the PRs will be merged as is, so PRs may be based on the final commit of the previous PR. The PRs should include README updates if functionality is added. Also please tag with completed (if you work with the GitHub web, you will need to create a release for this, which can be done from the repository main page).

# Lunch Menu API - Security Enhancement Assignment

The newest intern at the NBIS SysDev team started working with us recently and directly recognized the need of having an API that provides information about the lunch menu from the restaurants around the office. Therefore, proceeded to build a first version of the API, that would eventually be deployed on our servers and help the team to decide where to have lunch on each day.

## Overview

The current API provides public access to restaurant and menu data. You, as our new backend developer with expertise in security, are assigned the task of extending the existing lunch menu API to include user authentication, authorization, and to improve overall security. 

## Tasks

Specifically, the tasks that you are supposed to work on as part of this assignment are the following:

### Task 1

The lunch menu API is currently completely open and the endpoints can be accessed by anyone. The first task is to implement access control mechanisms, which would allow for an administrator to interact with the database and make changes to it. In this case, the information endpoints can remain open, while the ones related to administration should only be accessible by authorized people. It is up to you to decide the authentication and authorization mechanism that you will implement for this use case and what improvements you would like to make in the existing code.

### Task 2

The second task involves the implementation of the endpoints required in order to enable CRUD functionality for the administrators of the API. Specifically, an administrator should be able to: 
1. Create new entries in the database. 
2. Delete existing entries from the database.
3. Update existing entries in the database.

It is up to you to decide what information should be accessible. 

Please focus on the way the access is granted and on the secure implementation of the features, rather than on the number of functionalities. In other words, it is more important to have one endpoint that is safely adding data to the database, than multiple of them that lack security.

### (Optional) Task 3

The optional task is related to the deployment of your solution. For this task, create a kubernetes deployment setup for the application with proper configuration management and security considerations for production deployment. 

Please ensure detailed documentation of the deployment process is provided.

In case the task is too long, feel free to describe the considerations you would make, instead of providing the code itself.


## General instructions
Please provide not only working but acceptable code. In this case, that includes an appropriate level of documentation, testing, and so on. Assume that potential future colleagues will review your code and use it in upcoming projects.

You are not required to write your implementation from scratch, feel free to build on existing libraries, as well as specific frameworks.

It should be possible to deploy your final solution by following the instructions you include in the repository.

## Questions

If you have any questions or need any kind of feedback please open a new issue and tag the team @NBISweden/sysdev-recruitment when you add your question.


Good luck!
