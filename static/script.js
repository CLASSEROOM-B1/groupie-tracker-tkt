background=document.querySelector('.background')
bg=document.querySelector('.bg')
nom=document.querySelector('.nom')
user=document.querySelector('.user')

button.onclick = function nique() {
    nom.style.filter = 'blur(0px)';
    user.style.filter = 'blur(0px)';
    button.style.fontSize = '0px'
    button.style.height = '0px';
    button.style.width = '0px';
    setTimeout(pute, 250)
  };

  function pute() {
    button.style.display = 'none';
  }

  console.log("OK")