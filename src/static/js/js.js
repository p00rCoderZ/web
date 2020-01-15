function newAd(){
	$('#main').html(`
	<div id=info>Wypełnij poniższe pola aby dodać ogłoszenie</div>
	<div id=widePanel >

		<form action="add.html" method="post" id="create_post">
			Tytuł: <input type="text" name="title"><br>
			Treść: <textarea name="content" form="create_post"></textarea><br>
		
			<input type="radio" id="seek" name="type" value="Szukam pomocy">
			<label for="seek">Szukam pomocy</label>
		
			<input type="radio" id="offer" name="type" value="Oferuję pomoc">
			<label for="offer">Oferuję pomoc</label>
			<br>
		
			<select name="tags" size="3" multiple>
				<option value="tag1">tag1</option>
				<option value="tag2">tag2</option>
				<option value="tag3">tag3</option>
			</select>
		
		
			<br>
			<input type="submit" value="Prześlij">
		</form>
	</div>
	`);
	
}

function logIn(){
	$('#main').html(`
	<div id=info>Logowanie</div>
	<div id=widePanel >
		<form action='login.html' method='post' id='login_form'>
		    Email: <input type='text' name='email'><br>
		    Hasło: <input type='password' name='password'><br>

		    <input type='submit' value='Zaloguj'>
		</form>
	</div>
	`);
}

function signUp(){
	$('#main').html(`
	<div id=info>Rejestracja</div>
	<div id=widePanel >
		<form action='signup.html' method='post' id='signup_form'>

			<div class="form-group">
				<label for="exampleInputEmail1">Email</label>
				<input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Wpisz adres email">
				<small id="emailHelp" class="form-text text-muted">Twój adres e-mail zostanie z nami!</small>
			</div>
			
			
			<div class="form-group">			
				<label for="exampleInputLogin1">Login</label>
				<input type="text" class="form-control" id="exampleInputLogin1" placeholder="Login">
			</div>
			
			
			<div class="form-group">			
				<label for="exampleInputPassword1">Hasło</label>
				<input type="password" class="form-control" id="exampleInputPassword1" placeholder="Hasło">
			 </div>
			

			<button type="submit" class="btn btn-dark">Zarejestruj</button>
		</form>
	</div>


	`);
		
}

function mainPage(){
	location.reload(true);
}

