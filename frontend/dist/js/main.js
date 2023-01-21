const signup =document.getElementById('signup');
const login =document.getElementById('login');
const editUser =document.getElementById('edituser');
const editProduct =document.getElementById('editproduct');
const loadAdmin =document.getElementById('loadAdmin');

  signup.addEventListener('submit',signupFunc);
  login.addEventListener('submit',loginFunc);
  loadAdmin.addEventListener('click',loadAdminPage);

async function signupFunc(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const data = Array.from(formData.entries()).reduce((memo, [key, value]) => ({
      ...memo,
      [key]: value,
    }), {});
    console.log(data);
    const xhr = new XMLHttpRequest();
    let dataSt=JSON.stringify(data);
    console.log(dataSt);
    console.log(data.UserName);
    let user=data.UserName;
    $.ajax({
      type: 'post',
      url: 'http://localhost:4000/home/user',   
      data: data,
      xhrFields: {
          withCredentials: false
      },
      body:data
      ,
      success: function () {
        // login(user); 
        loadAdminPage(user);
        console.log(user);
      },
      error: function () {
          console.log('We are sorry but our servers are having an issue right now');
      }
    });
}
async function loadAdminPage(e) {
  e.preventDefault();
  console.log('Loading admin page');
  let xhr = new XMLHttpRequest();
  xhr.open('GET', '../pages/adminpanel.html', true);
  xhr.send()
}
async function loginFunc(e){
  const formData = new FormData(e.target);
  const data = Array.from(formData.entries()).reduce((memo, [key, value]) => ({
    ...memo,
    [key]: value,
  }), {});
  console.log(data);
  userName=data.UserName;
  password=data.password;
$.ajax({
  type: 'get',
  url:`http://localhost:4000/home/user/username=${userName}&password=${password}`,
  xhrFields: {
    withCredentials: false
},
success: function (){
},
error: function () { 
}
})
}