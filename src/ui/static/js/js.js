function newAd(){
	$('#main').html(`
	<div id=info>Wypełnij poniższe pola aby dodać ogłoszenie</div>
	<div id=widePanel3>
	<form>
		<div class="form-group" method="post">
			<label for="exampleTitle">Tytuł</label>
			<input type="text" class="form-control" id="exampleTitle" aria-describedby="emailHelp" placeholder="Podaj tytuł ogłoszenia">
		<div class="form-group">
			<label for="exampleContent">Treść</label>
			<textarea class="form-control" id="exampleContent" rows="3" placeholder="Podaj treść ogłoszenia"></textarea>
		</div>
	
		<div class="form-check">
			<input class="form-check-input" type="radio" name="exampleRadios" id="choice1" value="seek" checked>
			<label class="form-check-label" for="choice1">
				Szukam pomocy
			</label>
		</div>
		<div class="form-check">
			<input class="form-check-input" type="radio" name="exampleRadios" id="choice2" value="offer">
			<label class="form-check-label" for="choice2">
				Oferuję pomoc
			</label>
		</div>
		<br>
		<div class="form-group">
			<label for="exampleSelect">Wybierz tagi</label>
			<select multiple class="form-control" id="exampleSelect">
				<option>1</option>
				<option>2</option>
				<option>3</option>
				<option>4</option>
				<option>5</option>
			</select>
			<small id="selectHelp" class="form-text text-muted">Przytrzymaj ctrl aby wybrać kilka tagów</small>
		</div>
	
		<button type="submit" class="btn btn-dark">Dodaj</button>
	  </form>
	  </div>
	`);
	
}

function logIn(){
	$('#main').html(`
	<div id=info>Logowanie</div>
	<div id=widePanel2>
    <form>
        <div class="form-group" method="post">
            <label for="exampleInputEmail1">Email</label>
            <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Podaj adres email">
        </div>
        <div class="form-group">
            <label for="exampleInputPassword1">Hasło</label>
            <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Podaj hasło">
        </div>
        <button type="submit" class="btn btn-dark">Zaloguj</button>
	</form>
	

	</div>
	`);
}

function signUp(){
	$('#main').html(`
	<div id=info>Rejestracja</div>
	<div id=widePanel2>
		<form action='signup.html' method='post' id='signup_form'>

			<div class="form-group">
				<label for="exampleInputEmail1">Email</label>
				<input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Podaj adres email">
				<small id="emailHelp" class="form-text text-muted">Twój adres e-mail zostanie z nami!</small>
			</div>
			
			
			<div class="form-group">			
				<label for="exampleInputLogin1">Login</label>
				<input type="text" class="form-control" id="exampleInputLogin1" placeholder="Podaj login">
			</div>
			
			
			<div class="form-group">			
				<label for="exampleInputPassword1">Hasło</label>
				<input type="password" class="form-control" id="exampleInputPassword1" placeholder="Podaj hasło">
			 </div>
			

			<button type="submit" class="btn btn-dark">Zarejestruj</button>
		</form>
	</div>


	`);
		
}

function mainPage(){
	location.reload(true);
}
function advertisements(){
	$('#logo').html(`
	<style>
		nav{
			background-color:white;
			position:fixed;

			index-z:1000;
		}
		#info{
			margin-top:2em;
		}	
	</style>
	
	
	`);
	
	//Place holder \/
	$('#main').html(`
	<div id=info>Placeholder dla strony z ogłoszeniami</div>
	<div id=widePanel2>
		<form action='signup.html' method='post' id='signup_form'>

			<div class="form-group">
				<label for="exampleInputEmail1">Email</label>
				<input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Podaj adres email">
				<small id="emailHelp" class="form-text text-muted">Twój adres e-mail zostanie z nami!</small>
			</div>
			
			
			<div class="form-group">			
				<label for="exampleInputLogin1">Login</label>
				<input type="text" class="form-control" id="exampleInputLogin1" placeholder="Podaj login">
			</div>
			
			
			<div class="form-group">			
				<label for="exampleInputPassword1">Hasło</label>
				<input type="password" class="form-control" id="exampleInputPassword1" placeholder="Podaj hasło">
			 </div>
			

			<button type="submit" class="btn btn-dark">Zarejestruj</button>
		</form>
	</div>


	`);
	
}
