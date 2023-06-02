# Status Check

This is a simple single-page application (SPA) to check and display the status of certain URLs. It uses Go for the back-end API and Vite with React for the front-end. 

The application periodically checks the status of predefined URLs (currently www.google.com and www.amazon.com) and provides an interface for manually checking the status of these URLs. The status includes the HTTP status code, duration of the request, and the time of the check.

## Technologies used:

1. Go
2. React
3. Vite

## NPM Packages:

1. axios
2. react
3. react-dom
4. @vitejs/plugin-react
5. vite

## Installation instructions:

1. Clone the repository to your local machine.

2. Navigate to the project folder.

3. Install the npm packages:
    ```
    npm run install-all
    ```

4. Run the application and server:
    ```
    npm run start
    ```

5. If you want to run the API tests:
    ```
    npm run api-test
    ```

Please note that Go needs to be installed on your system to run the backend server and tests.

## Created by Bilal Khundmiri for Sun Life.
