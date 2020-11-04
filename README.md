# Go API client for gohiarc

Welcome to the Hiarc API documentation.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./gohiarc"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:5000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | [**InitDB**](docs/AdminApi.md#initdb) | **Post** /admin/database/init | Init the Hiarc database
*AdminApi* | [**ResetDB**](docs/AdminApi.md#resetdb) | **Put** /admin/database/reset | Reset the Hiarc database
*ClassificationApi* | [**CreateClassification**](docs/ClassificationApi.md#createclassification) | **Post** /classifications | Create a Classification
*ClassificationApi* | [**DeleteClassification**](docs/ClassificationApi.md#deleteclassification) | **Delete** /classifications/{key} | Delete a Classification
*ClassificationApi* | [**FindClassification**](docs/ClassificationApi.md#findclassification) | **Post** /classifications/find | Find a Classification
*ClassificationApi* | [**GetAllClassifications**](docs/ClassificationApi.md#getallclassifications) | **Get** /classifications | Get all Classifications
*ClassificationApi* | [**GetClassification**](docs/ClassificationApi.md#getclassification) | **Get** /classifications/{key} | Get a Classification&#39;s Info
*ClassificationApi* | [**UpdateClassification**](docs/ClassificationApi.md#updateclassification) | **Put** /classifications/{key} | Update a Classification
*CollectionApi* | [**AddChildToCollection**](docs/CollectionApi.md#addchildtocollection) | **Put** /collections/{key}/children/{childKey} | Add a child item to a Collection
*CollectionApi* | [**AddFileToCollection**](docs/CollectionApi.md#addfiletocollection) | **Put** /collections/{key}/files | Add a File to a Collection
*CollectionApi* | [**AddGroupToCollection**](docs/CollectionApi.md#addgrouptocollection) | **Put** /collections/{key}/groups | Add a Group to a Collection
*CollectionApi* | [**AddUserToCollection**](docs/CollectionApi.md#addusertocollection) | **Put** /collections/{key}/users | Add a User to a Collection
*CollectionApi* | [**CreateCollection**](docs/CollectionApi.md#createcollection) | **Post** /collections | Create a Collection
*CollectionApi* | [**DeleteCollection**](docs/CollectionApi.md#deletecollection) | **Delete** /collections/{key} | Delete a Collection
*CollectionApi* | [**FindCollection**](docs/CollectionApi.md#findcollection) | **Post** /collections/find | Find a Collection
*CollectionApi* | [**GetAllCollections**](docs/CollectionApi.md#getallcollections) | **Get** /collections | Get all Collections
*CollectionApi* | [**GetCollection**](docs/CollectionApi.md#getcollection) | **Get** /collections/{key} | Get a Collection&#39;s Info
*CollectionApi* | [**GetCollectionChildren**](docs/CollectionApi.md#getcollectionchildren) | **Get** /collections/{key}/children | Get a Collection&#39;s child Collections
*CollectionApi* | [**GetCollectionFiles**](docs/CollectionApi.md#getcollectionfiles) | **Get** /collections/{key}/files | Get a Collection&#39;s Files
*CollectionApi* | [**GetCollectionItems**](docs/CollectionApi.md#getcollectionitems) | **Get** /collections/{key}/items | Get a Collection&#39;s child items, including Collections and Files
*CollectionApi* | [**RemoveFileFromCollection**](docs/CollectionApi.md#removefilefromcollection) | **Delete** /collections/{key}/files/{fileKey} | Remove a File from a Collection
*CollectionApi* | [**UpdateCollection**](docs/CollectionApi.md#updatecollection) | **Put** /collections/{key} | Update a Collection
*FileApi* | [**AddClassificationToFile**](docs/FileApi.md#addclassificationtofile) | **Put** /files/{key}/classifications | Add a Classification to a File
*FileApi* | [**AddGroupToFile**](docs/FileApi.md#addgrouptofile) | **Put** /files/{key}/groups | Give a group access to a File
*FileApi* | [**AddRetentionPolicyToFile**](docs/FileApi.md#addretentionpolicytofile) | **Put** /files/{key}/retentionpolicies | Add a Retention Policy to a File
*FileApi* | [**AddUserToFile**](docs/FileApi.md#addusertofile) | **Put** /files/{key}/users | Give a user access to a File
*FileApi* | [**AddVersion**](docs/FileApi.md#addversion) | **Put** /files/{key}/versions | Add a new File Version
*FileApi* | [**AttachToExisitingFile**](docs/FileApi.md#attachtoexisitingfile) | **Put** /files/{key}/attach | Attach to an existing File
*FileApi* | [**CopyFile**](docs/FileApi.md#copyfile) | **Put** /files/{key}/copy | Copy a File
*FileApi* | [**CreateDirectUploadUrl**](docs/FileApi.md#createdirectuploadurl) | **Post** /files/directuploadurl | Create a direct upload url to a storage service
*FileApi* | [**CreateFile**](docs/FileApi.md#createfile) | **Post** /files | Create a File
*FileApi* | [**DeleteFile**](docs/FileApi.md#deletefile) | **Delete** /files/{key} | Delete a File
*FileApi* | [**DownloadFile**](docs/FileApi.md#downloadfile) | **Get** /files/{key}/download | Download a File
*FileApi* | [**GetCollectionsForFile**](docs/FileApi.md#getcollectionsforfile) | **Get** /files/{key}/collections | Get a list of Collections for a File
*FileApi* | [**GetDirectDownloadUrl**](docs/FileApi.md#getdirectdownloadurl) | **Get** /files/{key}/directdownloadurl | Get a direct download URL to a File
*FileApi* | [**GetFile**](docs/FileApi.md#getfile) | **Get** /files/{key} | Get a File&#39;s Info
*FileApi* | [**GetRetentionPolicies**](docs/FileApi.md#getretentionpolicies) | **Get** /files/{key}/retentionpolicies | Get a list of Retention Policies on a File
*FileApi* | [**GetVersions**](docs/FileApi.md#getversions) | **Get** /files/{key}/versions | Get a list of File Versions
*FileApi* | [**UpdateFile**](docs/FileApi.md#updatefile) | **Put** /files/{key} | Update a File
*FilesApi* | [**FilterAllowedFiles**](docs/FilesApi.md#filterallowedfiles) | **Post** /files/allowed | Filter a list of File keys that a User can access
*GroupApi* | [**AddUserToGroup**](docs/GroupApi.md#addusertogroup) | **Put** /groups/{key}/users/{userKey} | Add a User to a Group
*GroupApi* | [**CreateGroup**](docs/GroupApi.md#creategroup) | **Post** /groups | Create a Group
*GroupApi* | [**DeleteGroup**](docs/GroupApi.md#deletegroup) | **Delete** /groups/{key} | Delete a Group
*GroupApi* | [**FindGroup**](docs/GroupApi.md#findgroup) | **Post** /groups/find | Find a Group
*GroupApi* | [**GetAllGroups**](docs/GroupApi.md#getallgroups) | **Get** /groups | Get all Groups
*GroupApi* | [**GetGroup**](docs/GroupApi.md#getgroup) | **Get** /groups/{key} | Get a Group&#39;s Info
*GroupApi* | [**GetGroupsForCurrentUser**](docs/GroupApi.md#getgroupsforcurrentuser) | **Get** /users/current/groups | Get the Groups for the current User
*GroupApi* | [**UpdateGroup**](docs/GroupApi.md#updategroup) | **Put** /groups/{key} | Update a Group
*GroupsApi* | [**GetGroupsForUser**](docs/GroupsApi.md#getgroupsforuser) | **Get** /users/{key}/groups | Get Groups for a User
*LegalHoldApi* | [**CreateLegalHold**](docs/LegalHoldApi.md#createlegalhold) | **Post** /legalholds | Create a Legal Hold
*LegalHoldApi* | [**GetLegalHold**](docs/LegalHoldApi.md#getlegalhold) | **Get** /legalholds/{key} | Get a Legal Hold&#39;s Info
*RetentionPolicyApi* | [**CreateRetentionPolicy**](docs/RetentionPolicyApi.md#createretentionpolicy) | **Post** /retentionpolicies | Create a Retention Policy
*RetentionPolicyApi* | [**FindRetentionPolicies**](docs/RetentionPolicyApi.md#findretentionpolicies) | **Post** /retentionpolicies/find | Find a Retention Policy
*RetentionPolicyApi* | [**GetAllRetentionPolicies**](docs/RetentionPolicyApi.md#getallretentionpolicies) | **Get** /retentionpolicies | Get all Retention Policies
*RetentionPolicyApi* | [**GetRetentionPolicy**](docs/RetentionPolicyApi.md#getretentionpolicy) | **Get** /retentionpolicies/{key} | Get a Retention Policy&#39;s Info
*RetentionPolicyApi* | [**UpdateRetentionPolicy**](docs/RetentionPolicyApi.md#updateretentionpolicy) | **Put** /retentionpolicies/{key} | Update a Retention Policy
*TokenApi* | [**CreateUserToken**](docs/TokenApi.md#createusertoken) | **Post** /tokens/user | Create a Token for a User
*UserApi* | [**CreateUser**](docs/UserApi.md#createuser) | **Post** /users | Create a User
*UserApi* | [**DeleteUser**](docs/UserApi.md#deleteuser) | **Delete** /users/{key} | Delete a User
*UserApi* | [**FindUser**](docs/UserApi.md#finduser) | **Post** /users/find | Find a User
*UserApi* | [**GetAllUsers**](docs/UserApi.md#getallusers) | **Get** /users | Get all Users
*UserApi* | [**GetCurrentUser**](docs/UserApi.md#getcurrentuser) | **Get** /users/current | Get the current User
*UserApi* | [**GetGroupsForCurrentUser**](docs/UserApi.md#getgroupsforcurrentuser) | **Get** /users/current/groups | Get the Groups for the current User
*UserApi* | [**GetGroupsForUser**](docs/UserApi.md#getgroupsforuser) | **Get** /users/{key}/groups | Get Groups for a User
*UserApi* | [**GetUser**](docs/UserApi.md#getuser) | **Get** /users/{key} | Get a User
*UserApi* | [**UpdateUser**](docs/UserApi.md#updateuser) | **Put** /users/{key} | Update a User


## Documentation For Models

 - [AccessLevel](docs/AccessLevel.md)
 - [AddClassificationToFileRequest](docs/AddClassificationToFileRequest.md)
 - [AddFileToCollectionRequest](docs/AddFileToCollectionRequest.md)
 - [AddGroupToCollectionRequest](docs/AddGroupToCollectionRequest.md)
 - [AddGroupToFileRequest](docs/AddGroupToFileRequest.md)
 - [AddRetentionPolicyToFileRequest](docs/AddRetentionPolicyToFileRequest.md)
 - [AddUserToCollectionRequest](docs/AddUserToCollectionRequest.md)
 - [AddUserToFileRequest](docs/AddUserToFileRequest.md)
 - [AddVersionToFileRequest](docs/AddVersionToFileRequest.md)
 - [AllowedFilesRequest](docs/AllowedFilesRequest.md)
 - [AttachToExistingFileRequest](docs/AttachToExistingFileRequest.md)
 - [Classification](docs/Classification.md)
 - [Collection](docs/Collection.md)
 - [CollectionItems](docs/CollectionItems.md)
 - [CopyFileRequest](docs/CopyFileRequest.md)
 - [CreateClassificationRequest](docs/CreateClassificationRequest.md)
 - [CreateCollectionRequest](docs/CreateCollectionRequest.md)
 - [CreateDirectUploadUrlRequest](docs/CreateDirectUploadUrlRequest.md)
 - [CreateFileRequest](docs/CreateFileRequest.md)
 - [CreateFileRequestAllOf](docs/CreateFileRequestAllOf.md)
 - [CreateGroupRequest](docs/CreateGroupRequest.md)
 - [CreateLegalHoldRequest](docs/CreateLegalHoldRequest.md)
 - [CreateOrUpdateEntityRequest](docs/CreateOrUpdateEntityRequest.md)
 - [CreateRetentionPolicyRequest](docs/CreateRetentionPolicyRequest.md)
 - [CreateRetentionPolicyRequestAllOf](docs/CreateRetentionPolicyRequestAllOf.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateUserTokenRequest](docs/CreateUserTokenRequest.md)
 - [Entity](docs/Entity.md)
 - [File](docs/File.md)
 - [FileAllOf](docs/FileAllOf.md)
 - [FileDirectDownload](docs/FileDirectDownload.md)
 - [FileDirectUpload](docs/FileDirectUpload.md)
 - [FileVersion](docs/FileVersion.md)
 - [FindClassificationsRequest](docs/FindClassificationsRequest.md)
 - [FindCollectionsRequest](docs/FindCollectionsRequest.md)
 - [FindEntityRequest](docs/FindEntityRequest.md)
 - [FindGroupsRequest](docs/FindGroupsRequest.md)
 - [FindRetentionPoliciesRequest](docs/FindRetentionPoliciesRequest.md)
 - [FindUsersRequest](docs/FindUsersRequest.md)
 - [Group](docs/Group.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [LegalHold](docs/LegalHold.md)
 - [RetentionPolicy](docs/RetentionPolicy.md)
 - [RetentionPolicyApplication](docs/RetentionPolicyApplication.md)
 - [UpdateClassificationRequest](docs/UpdateClassificationRequest.md)
 - [UpdateCollectionRequest](docs/UpdateCollectionRequest.md)
 - [UpdateFileRequest](docs/UpdateFileRequest.md)
 - [UpdateGroupRequest](docs/UpdateGroupRequest.md)
 - [UpdateRetentionPolicyRequest](docs/UpdateRetentionPolicyRequest.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [User](docs/User.md)
 - [UserCredentials](docs/UserCredentials.md)


## Documentation For Authorization



## AdminApiKeyAuth

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```


## JWTBearerAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```



## Author


