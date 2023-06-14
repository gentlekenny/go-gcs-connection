# Connect to Google Cloud Storage database using Go

I had troubles establishing connection to MySQL database which is live on Google Storage.
Documentation wasn't quite intuitive but I've figured out a way and I am sharing it with you.
To connect to GCS, you need a proxy dialer for mysql, which can be obtained from [Google Cloud Platform](github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql).
I also converted this connection into GORM, which I needed for my case. You can remove that part from the **ConnectToGCSMySQLDatabase** method, if you don't need it. Remember, you will also have to change return object from *gorm.DB* to *sql.DB*

**Please if this ever helps you, leave a star.** 