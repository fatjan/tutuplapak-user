# tutuplapak
Go backend app for the 3rd project in Projectsprint

# How to Run
1. Clone the repo

   ```bash
    git clone [git@github.com:fatjan/fitbyte.git](https://github.com/fatjan/tutuplapak-user)
    cd fitbyte
   ```

2. Create `.env` file

  Can copy from `.env-example` but adjust the value
   ```bash
    cp .env-example .env
   ```

3. Create database `tutuplapak`

4. Run the migration

  ```bash
    make migrate-up
  ```

5. Run the app

  ```bash
  make run
  ```
