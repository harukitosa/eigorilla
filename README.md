# eigorilla

## User Experience
1. Users can meet people who study English.
2. Talk with them and Practice use English.

## Value
1. simple and easy.
2. make English learning a habit.

## Function

must
1. login/logout
2. post and comment and view my posts 

should

3. mix conversation examples from time to time.


Not JWT token.


### Front 
* Vue.js or React

* Material Design

### Back
* Golang

### Auth
* Firebase

## URL structure FrontEnd


```
/ #landing page

/timeline #view someone's posts

/home #view your post

/login #login your account

/logout #logout your account

/comment #comment someone's post
```

## URL structure BackEnd

```
/post:{userID}  #User post.

/getTimeline #User get timeline date.

/getUserPost:{userID} #User get their own post.

/alert:{userID} #User get alert date.
```