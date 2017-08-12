This script takes input from stdin and uploads it to Google Cloud storage.

# Dependencies

go version `1.8`

# Installation

## Go environment

* Dowload go `1.8.3` from golang.org and extract if under `/usr/local` with 

```
tar -C /usr/local -xzf go1.8.3.$OS-$ARCH.tar.g
```
then link the bin executable 

```
ln -s /usr/local/go/bin/go /usr/local/bin
```

## Get dependencies and build

* Download `storage-upload.go`

```
wget https://raw.githubusercontent.com/nkcr/Google-Storage-stdin-Upload/master/storage-upload.go
```
 
and run `$ go get` within its parent folder.
* finally run `go build storage-upload.go` to build the executable.

# Usage

The script takes as input the standard input **stdin**. The most convenient way is to use pipe.

The script takes the followings parameters:

```
  -d string
        Upload path, of form 'gs://bucket/folder/file.txt'
  -help
        Prints usage
  -k string
        Service key path
  -p string
        Create the bucket in the corresponding given projectID if it doesn't exist
  -verbose
        Outputs infos
```

**-d** and **-k** are always required. For the key, see next section *security*.

Here is an example:

```
echo "Hello" | upload-sorage -k adfe43d.json -d gs://my_bucker/my_folder/file.txt
```

# Security

The script asks for a **service account** key with the params `-k`. You can create a service account key via your console.cloud.google.com on the section **IAM & Admin**. Then click on **Service accounts**. There you can create a service account key and assigning roles to it. Roles needed are **Storage/Storage Object Creator** if you only want to create files without override, **Storage/Storage Object Admin** if you want to override files or **Storage/Storage Admin** if you want to create new buckets (via `-p` params). Then don't forget to check **Furnish a new private key** in JSON format. This is the key you will need.