function newAd(){
	$('#main').html("<div id=info>Wypełnij poniższe pola aby dodać ogłoszenie!</div><div id=widePanel ><form method='post' id='create_post'>    Tytuł: <input type='text' name='title'><br>    Treść: <textarea name='content' form='create_post'></textarea><br>    <input type='radio' id='seek' name='type' value='Szukam pomocy'>    <label for='seek'>Szukam pomocy</label>    <input type='radio' id='offer' name='type' value='Oferuję pomoc'>    <label for='offer'>Oferuję pomoc</label>    <br>    <select name='tags' size='3' multiple>        <option value='tag1'>tag1</option>        <option value='tag2'>tag2</option>        <option value='tag3'>tag3</option   </select>    <br>    <input type='submit' value='Prześlij'>  </form></div>");	
	
}

function logIn(){
	$('#main').html("<div id=info>Logowanie</div><div id=widePanel ><form action='login.html' method='post' id='login_form'>Login: <input type='text' name='login'><br>Hasło: <input type='password' name='password'><br><input type='submit' value='Zaloguj'></form></div>");
}

function mainPage(){
	location.reload(true);
}