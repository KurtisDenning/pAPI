# Admin App

This is a simple application that was developed to make it easier to interact with our data. It does this by automatically connecting to our pAPI database located in MongoDB - all that it requires is the user to enter the password for the `db_admin` user. From this application, all the CRUD operations that are relevant to our database are allowed, with some simple error checking to minimise human error during data entry.

***Note:** This application is for administrator use only. As such, there are some quirks that might make it difficult to use, but as an admin you should be familiar with the application anyway, and so this shouldn't be a problem. This application was developed as an extra application to make it easy to add data to the database, and was an afterthought to this project. It is **not** for public use as it allows connecting directly to the database.*

## Installation + Connection

### **Linux**

1. Download the Linux binary [here](https://github.com/oliverschweikert/pAPI/raw/main/admin-app/builds/papiAdmin-linux)
2. Make the file executable - `sudo chmod u+x papiAdmin-linux` - or Right Click > Permissions > Allow executing file as a program.

    ![Allow executing file as program](docs/images/walkthrough04.png)
3. Double click to launch the application.
4. Click `Connect`

    ![Connect](docs/images/walkthrough05.png)
5. Enter the password (supplied via email :P) and click connect.

    ![Enter password](docs/images/walkthrough06.png)
6. Well done :)

### Windows

1. Download and run the windows .exe [here](https://github.com/oliverschweikert/pAPI/raw/main/admin-app/builds/papiAdmin-windows.exe).
2. Because it's an unsigned application, Windows may block it to protect you. Feel free to ignore the warnings, the application does nothing but connect to the database, you can view the code if you don't believe me ;)

## CRUD Operations

There are numerous Create, Read, Update and Delete operations that this application makes availale. There is basic data validation implemented by default on these operations to stop the user from making a change that would result in corrupt data that won't map properly to the models defined in the backend application.

### **Categories**

Following are the CRUD operations for the `Categories` collection. To perform these operations, first enter the `Categories` window.

    ![Categories Window](docs/images/walkthrough07.png)

#### **Create Category**

There is a button in the `category` window that provides functionality to add a new category. This is a form that allows entry for three things:

- Name
- Short Description
- Long Description

The `ID` field is automatically generated for the user.

Steps to perform this operation:

1. Click the button `New Category`.

    ![New Category](docs/images/walkthrough08.png)
2. Enter new category data and click submit.

    ![Enter data](docs/images/walkthrough09.png)
3. If it was added to the database, you will see a `Success` dialog. Clock `ok`.

    ![Success](docs/images/walkthrough10.png)
4. Close the current categories window, and reopen it to get the up to date information.

    ![Updated categories](docs/images/walkthrough11.png)

#### **Read Category**

All categories are displayed as accordion items in the `categories` window. The fields that are displayed are:

- ID
- Name
- Short Description
- Long Description

Steps to perform this operation:

1. Just expand an accordion to get detailed information about the item in the database.

    ![Details](docs/images/walkthrough12.png)

#### **Update Category**

Be careful when modifying data, you can mess up the appication!

The categories displayed in the `categories` window are data entry fields. These can be modified, and when the update button is clicked, these changes will be applied. The fields displayed are the same as the [Read Category](#read-category) section, but the fields that are actually modified on update are the same as the [Create Category](#create-category) section. This is because Mongo doesn't allow modifying the ID field.

Steps to perform this operation:

1. Change some of the fields in the details section of the category. Note that while you can change all the text fields, only some will actually update. For example, even if you change the ID, it won't be updated for database consistency.

    ![Update fields](docs/images/walkthrough13.png)
2. Push `Update`.

    ![Update](docs/images/walkthrough14.png)
3. Press `Ok` if successful.

    ![Success](docs/images/walkthrough15.png)

#### **Delete Category**

Be careful when modifying data, you can mess up the application!

A category can be removed by pressing the `Delete` button inside a category accordion. The result is simple - the category is deleted from the database.

Steps to perform this operation:

1. Press delete in the details section of a category.

    ![Delete](docs/images/walkthrough16.png)
2. Confirm the deletion.

    ![Confirm](docs/images/walkthrough17.png)
3. If successful, the `Categories` window will automatically close. Reopen the window to view the changes.

    ![Reopen](docs/images/walkthrough18.png)

### **API Data**

Following are the CRUD operations for the `apiData` collection.

Steps to open API Data:

1. Click `Info` for API data

    ![Info](docs/images/walkthrough19.png)

#### **Create API Data**

There is a button in the `API Data` window that provides functionality to add a new category. This is a form that allows entry for three things:

- Title
- Description
- External URL
- Base URL

The following fields are automatically generated for the user:

- ID
- Categories
- Daily Count
- Weekly Count
- Monthly Count
- Yearly Count
- Total Count
- Requests Array
- Categories Array

There is also the ability to add new `Requests` and `Categories` to an individual set of `API Data`, since these are arrays. To add a new request, open the `Requests` accordion and click the button labelled **New Request**. The available fields for a new request are:

- URL

These request fields are automatically generated for the user:

- Last Update
- Response

To add a new category, open the `Categories` accordion and click **New Category**. The available fields for a new category are:

- Category Hex

Please note that the hex value must match the hex value of a category stored in the database, or an error will be shown.

Steps to add a new API Data:

1. Click `New API Data`.

    ![New API Data](docs/images/walkthrough20.png)
2. Enter some API data and click submit.

    ![Submit](docs/images/walkthrough21.png)
3. If successful, click `Ok`.

    ![Success](docs/images/walkthrough22.png)
4. Close the `API Data` window and reopen it to see updated data.

    ![Update](docs/images/walkthrough23.png)

Steps to add a new request:

1. Expand a set of API data, click `Requests` and then `Add request`.

    ![Add Request](docs/images/walkthrough24.png)
2. Enter the sub-url to add on to the base URL, and click `Add`.

    ![Add](docs/images/walkthrough25.png)
3. Is successful, click `Ok` on the dialog

    ![Ok](docs/images/walkthrough26.png)
4. Close the `API Data` window and reopen it to refresh the data.

   ![Refresh](docs/images/walkthrough27.png)

Steps to add a new category:

1. Expand a set of API data, click `Categories` and then `Add category`.

    ![Add Category](docs/images/walkthrough28.png)
2. Enter hex of a category, and click `Add`.

    ![Add](docs/images/walkthrough29.png)
3. Is successful, click `Ok` on the dialog

    ![Ok](docs/images/walkthrough30.png)
4. Close the `API Data` window and reopen it to refresh the data.

   ![Refresh](docs/images/walkthrough31.png)

#### **Read API Data**

All API data is displayed as accordion items in the `API Data` window. The fields that are displayed are:

- ID
- Title
- Description
- External URL
- Base URL
- Requests (Accordion)
- Categories (Accordion)

To see all the information, just open up all the accordions:

    ![Read](docs/images/walkthrough32.png)

#### **Update API Data**

Be careful when modifying data, you can mess up the appication!

The API data displayed in the `API Data` window are data entry fields. These can be modified, and when the update button in the main window is clicked, these changes will be applied. The fields displayed are the same as the [Read API Data](#read-api-data) section, but the fields that are actually modified on update are the same as the form fields in the [Create API Data](#create-api-data) section. This is because Mongo doesn't allow modifying the ID field, and these other values are calculated based off client interaction with the site.

Steps to perform an update on APIData:

1. Modify one of the detail fields and click `Update`. Note that you cannot modify a `request` or a `category`, you will need to delete it and add a new one.

   ![Modify details](docs/images/walkthrough33.png)
2. If successful, press `Ok` on the dialog.

   ![Update](docs/images/walkthrough34.png)

#### **Delete API Data**

Be careful when modifying data, you can mess up the application!

A set of API data can be removed by pressing the `Delete` button inside an API Data accordion. The result is simple - the set of API data is deleted from the database. Be careful with this, it's a pain to enter all the necessary information again if you accidentally delete something.

The steps to delete API data are:

1. Push the delete button within the details.

    ![Delete API data](docs/images/walkthrough37.png)
2. Press `Yes` to confirm deletion.

    ![Confirm](docs/images/walkthrough38.png)
3. If successful, press `Ok` on the dialog to automatically close the API data.

    ![Ok](docs/images/walkthrough39.png)
4. Reopen the window to view the updated data

   ![Reopen](docs/images/walkthrough40.png)

To delete an API category, press `Remove` on the category, press `Delete` to confirm, and then close and reopen the API data window to view the updated data.

![Delete category](docs/images/walkthrough35.png)

To delete an API request, press `Remove` on the request, press `Delete` to confirm, and then close and reopen the API data window to view the updated data.

![Delete request](docs/images/walkthrough36.png)

## Technologies

This application was developed using [Fyne](https://fyne.io/), a GUI framework for [Go](https://go.dev/). This was chosen because it is quick to develop a simple, offline GUI app, similar to Windows Forms in C#. We chose to develop the application offline, and run it as a desktop app for security, since the application provides admin level access to a database. Also, because the database shouldn't require constant administration, it isn't necessary for this admin application to be running 24/7, and can just be started up when a modification is necessary.

## Future Maintenance

Currently, the functionality of the application should all be fine, but when making a change to the database, the windows don't automatically refresh with the most up to date information. A change that would improve the quality of life would be implementing the automatic refresh of the window when a change is detected. However, this isn't a necessary feature, since the target user of this application is an administrator, not a customer, and therefore the application doesn't necessarily need a positive user experience, as long as the necessary functionality is there.

## Firewall Considerations

This application directly connects to MongoDB. During development there were issues accessing the database through PEAP protected nectworks, such as Eduroam. If you are having issues connecting, please try and use an external network, such as mobile data. If the error continues persisting, contact the app developer.
