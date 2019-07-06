## Requirements 

* mandatory endpoints: create, read, update, delete (feel free to extend these endpoints) 

* use an authentication method to secure your requests (Examples: JWT token, oAuth, etc.) 

* use a way to make your data persistent (database is preferred) 

* write at least one unit test and a functional test  

## Endpoints 
*	**Read**: https://us-central1-johnbalvin.cloudfunctions.net/phonebook_read
 * Parameters(optional):  
       *  **id**(string) the id of the user
	  * Returning value:
		   * **Content type**: application/json
		   * **Body**: JSON representation of the user: id,name,address,phone_number
 * Without parameters:
		 * Returning value:
		   * **Content type**: application/json
		   * **Body**: JSON representation of all users users: id,name,address,phone_number up to 100 contacts
	   
*	**Create**: https://us-central1-johnbalvin.cloudfunctions.net/phonebook_create

  * **oAuth token**: ldfg08dsoas
  * Parameters:  
       *  **name**(string) the name of the user
	   *  **address**(string) the address of the user
	   *  **telephone**(string) the telephone of the user
  * Returning value:
       * **Content type**: application/json
	   * **Body**: JSON representation of the user: id,name,address,phone_number
	   
*	**Update**: https://us-central1-johnbalvin.cloudfunctions.net/phonebook_update

  * **oAuth token**: kfgjyi79
  * Parameters:  
       *  **field**(string) what field you want to update (telephone,name,address)
	   *  **value**(string) new value assign to that field
  * Returning value:
       *  NONE

*	**Delete**: https://us-central1-johnbalvin.cloudfunctions.net/phonebook_delete

  * **oAuth token**: kdsfpojui43
  * Parameters:  
       *  **id**(string) if from the contact you want to delete
  * Returning value:
       *  NONE
	      
## Information

* Endpoints for updating, creating and deleting need to send the header "Authorization: Bearer [token]" where [token] is the token for each endpoint
* Endpoint for reading does not need to send an authorization header
* In case you want to test the code locally you need to:
 1.  Create a [Google Cloud Account](https://console.cloud.google.com),   
 2.  Create a project,
 3.  Change projectID variable at phonebook/init.go by your projectID
 4.  Get credentials with [Google Cloud SDK](https://cloud.google.com/sdk)
 5. Run `go test -run TestCreateOK ` and `go test -run TestCreateFail`
 6. It should pass Go test
