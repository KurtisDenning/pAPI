# pAPI Front-end App

The purpose of the pAPI front-end is to provide a modern, user friendly interface to sort and display the information being stored in the back-end.

## Technologies

- React
- ES6
- Chakra-UI
- Npm

## Design

[Figma](https://www.figma.com/) was used to disign our front-end application which can be found [here](https://www.figma.com/file/jvhjc2HCBYN8FFp2c0OEwT/pAPI?node-id=0%3A1).

### UI

There were some careful design choices we took into consideration while designing the user interface such as its overall aesthetic, readability as well as how well it would scale when viewed across a wide range of devices.

Our initial prototype did have the functionality that we wanted from a front-end point of view but did not meet our standards from an aesthetic and usability stand point so a redesign was required. The original design as well as our redesign are both included in the Figma [file](https://https://www.figma.com/file/jvhjc2HCBYN8FFp2c0OEwT/pAPI?node-id=0%3A1).

pAPI was designed and built with a 'Mobile first' approach as there is currently a massive mobile demographic that will only continue to grow with time. Designing for mobile first means that our app not only works great on smaller mobile devices but will also scale to most if not all screen sizes whether they be mobile or not.

Thankfully the Chakra-UI library makes building a scalable app very easy with properties such as [Breakpoints](https://chakra-ui.com/docs/styled-system/theming/theme#breakpoints) and components like [Grid](https://chakra-ui.com/docs/components/layout/grid) to make building a scalable app quick and easy.

With all this in consideration, we decided to go with a very minimalistic design that is comprised of a mostly black and white color scheme which is promoted by an abundance of white space to provide great contrast between text and the background, aiding in the readability and usability of our app.

### UX

One of our goals that we set out to meet was to make our data fast and easily accessible to every user. This is why when designing our front end it was very important to us that it required as few clicks as possible to get the requested data.

Some of our UX features we integrated into our app were:

- A search bar.
  - This search bar is the very first thing that a user sees when visiting pAPI and is auto searching with state as the user types so when they hit the search button there is no need to wait for the query of that search to load.

- Category filtering.
  - The category filtering buttons are the very next thing the user sees and work in conjunction with the search bar to narrow down a users search even more if the user wishes to do so.

- Animations.
  - Some subtle animations are used throughout pAPI such as some scrolling animations as well as some loading animations which really add to the QOL of pAPI.
  
- Large UI elements
  - As talked about in the UI section, mobile has become a very large demographic of users so to accommodate those using smaller touch screens such as phones, we increased the size of all our functional UI elements such as buttons and links.

These features along with our simplistic design make the usability and speed of our app exceptional.

## Scalability

Since we are offering an admin app to developers so that they can easily add their API's to pAPI, it is crucial that our front-end can handle as much data as thrown at it.

Some features added to achieve this were:

- Pagination.
  - On the home page where the API's are being displayed, if the data is not filtered there is likely many items being displayed so it is important to have pagination to restrict how much data is displayed at one time.
  
- Logic to handle a variety different API data
  - With so many different unique API's, not all the data is formatted the same and some may even have images that need to be displayed. This is where that front-end logic comes into play, ensuring all data being displayed looks as best as it can.

## Testing

Given the time frame, we completed the entire barge of testing throughout development and then ran further tests on the application as a whole after the completion.

This ensures that each individual component is working as intended, and behaves properly when working with together with the rest of the site while saving us precious time for development that otherwise would have been spent on testing.

## Deployment

Currently our app is hosted on GitHub pages and can be viewed [here](https://oliverschweikert.github.io/pAPI).

For a real world deployment we would look into hosting the front-end component of our app on some reputable cloud hosting service such as Azure or AWS.
