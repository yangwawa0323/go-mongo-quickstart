# mongoDB go driver quick start

All the examples are put into **`my_testing`** folder.

You need create a environment setup file **`.env`** in the folder,

and set the MONGODB_URI variable to run the test. For example:


```shell
MONGODB_URI="mongodb+srv://<username><password>@cluster0.ct3kn.mongodb.net/<dbname>?retryWrites=true&w=majority"
```

> Note: replace the `username`, `password` and `dbname` according your database server