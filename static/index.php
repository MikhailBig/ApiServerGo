<?php
include 'adm_stuff/db.php';
?>
 <!DOCTYPE html>
 <html lang="en">
 <head>
   <meta charset="UTF-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <meta http-equiv="X-UA-Compatible" content="ie=edge">
   <?php
   $website_title = 'Bigun.Blog';
   include 'content/head.php';
   ?>
 </head>
 <body>
    <?php include 'content/preloader.php';?>
    <div class="wrapper">
      <div class="main">
        <?php include 'content/header.php';?>
        <div class="recent">
          <div class="grid grid-pad">
            <div class="col-1-8">
              <div class="content">
                <h3>Recent post:</h3>
              </div>
            </div>
            <div class="col-7-8"></div>
          </div>
        </div>
        <div class="main_content grid grid-pad">
          <div class="col-1-1 ">
            <div class="content">
              <h2>Sample Text</h2>
              <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <?php include 'content/footer.php';?>
</body>
</html>
