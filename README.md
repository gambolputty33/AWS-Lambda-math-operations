# Performing math operations with AWS Lambda, Golang, and JSON input/output.
## 1.  Signup for an AWS account.

https://www.youtube.com/watch?v=ne8LrbCzW0Q

## 2. Login as the root user and setup two-factor authentication using at least two Yubikeys.

<img width="365" height="680" alt="02" src="https://github.com/user-attachments/assets/fa7b202c-6718-433e-8018-170f84eec0a7" />

<img width="1803" height="896" alt="02b" src="https://github.com/user-attachments/assets/c6e53ef4-5d43-4f87-81ad-13e4891d7414" />

## 3. Create a non-root admin user, assign Administrator priviledges using the AWS IAM Identity Center, setup two-factor authentication using at least two Yubikeys, and get the AWS access portal URL.

https://www.youtube.com/watch?v=_KhrGFV_Npw

## 4. Login as the non-root user.

<img width="392" height="417" alt="04" src="https://github.com/user-attachments/assets/0dea12af-fff0-4da9-b4f0-38beb6f34151" />

## 5.  Select "AdministratorAccess" if you need to perform manual tasks within AWS.  Select "Access keys" to copy the AWS environment variables to the clipboard.

<img width="1805" height="390" alt="05" src="https://github.com/user-attachments/assets/5bdba61d-b321-43be-a812-f33008e72698" />

<img width="817" height="918" alt="05b" src="https://github.com/user-attachments/assets/4e3e397d-9a83-4a62-ba67-79e931b6c309" />

## 6. Install Visual Studio Code and Terraform on your PC.

## 7. Paste the AWS environment variables into your terminal window in Visual Studio Code.

## 8. Create a project directory in Visual Studio Code and copy the files from this repository into it.

## 9. Use the commands from the notes.txt file to prepare a Zip file for Terraform.

## 10. Use the "terraform init", "terraform plan", and "terraform apply" commands to deploy the Golang program.

<img width="1239" height="244" alt="08" src="https://github.com/user-attachments/assets/6f0ec700-9154-41d8-9a11-fce5df2157a2" />

## 11. Go to the Functions section of AWS Lambda to find the newly created Golang function.

<img width="1159" height="584" alt="09 " src="https://github.com/user-attachments/assets/03816ec8-1f52-4f08-bb13-e19b2baeaa75" />

## 12. Click on the Test icon.

<img width="1698" height="774" alt="10" src="https://github.com/user-attachments/assets/c4040515-cc56-4d15-80fa-d1d6fc4d786b" />

## 13. Scroll down to the "Event JSON" section, type in a name/value pair that the Golang program will accept as a properly formatted input, then click the Test icon.

<img width="1698" height="774" alt="13" src="https://github.com/user-attachments/assets/161a12b1-1532-4bfa-a70f-acab960ac8f5" />

## 14. The test results will show four math opertions on the two inputted numbers outputted in a JSON format.  Click on the "logs" link.  This will take you to Cloudwatch.

<img width="691" height="363" alt="14" src="https://github.com/user-attachments/assets/9582915d-23e3-41ea-92d7-6b0d3d4b7fab" />

## 15. In Cloudwatch, make a note of all of the log entries shown.

<img width="1802" height="450" alt="15" src="https://github.com/user-attachments/assets/b9b1fd65-5c59-4bb0-85f7-4a37d9c03fe7" />

## 16. Delete the Cloudwatch logs if you need to.

<img width="1802" height="450" alt="16" src="https://github.com/user-attachments/assets/b8b3e684-54ba-407a-900d-c5273ea42191" />

## 17. Use the "terraform destroy" command to delete the Golang program.

<img width="1262" height="280" alt="17" src="https://github.com/user-attachments/assets/68701e67-e52a-4a1d-a84e-a16c86e9f6d4" />
