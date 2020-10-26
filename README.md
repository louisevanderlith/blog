# blog
Blog.API - Stores all Blog and News articles
## Run with Docker
* $ docker build -t avosa/blog:dev .
* $ docker rm Blog
* $ docker run -d --network host --name Blog avosa/blog:dev
* $ docker logs Blog