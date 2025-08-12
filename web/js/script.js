function submitButton(email, password) {
  const url = 'http://localhost:8080/user';
  fetch(url, {
    method: 'POST',
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      email: email,
      password: password
    })
  })
  .then(res => res.json())
  .then(data => console.log("Resposta da API (Cadastro):", data))
  .catch(err => console.error("Erro ao cadastrar:", err));
}

const registerForm = document.getElementById("registerForm");
if (registerForm) {
  registerForm.addEventListener("submit", function(e) {
    e.preventDefault();
    const email = document.getElementById("registerEmail").value;
    const password = document.getElementById("registerPassword").value;
    submitButton(email, password);
  });
}

function submitButtonLogin(email, password) {
  const url = 'http://localhost:8080/login';
  fetch(url, {
    method: 'POST',
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password })
  })
  .then(res => {
    if (!res.ok) throw new Error(`Erro HTTP: ${res.status}`);
    return res.json();
  })
  .then(data => {
    console.log("Resposta da API (Login):", data);
    if (data.success) {
      alert("Usuário ou senha inválidos!");
    } else {
      alert("Login realizado com sucesso!");
    }
  })
  .catch(err => console.error("Erro ao fazer login:", err));
}

const loginForm = document.getElementById("login-form");
if (loginForm) {
  loginForm.addEventListener("submit", function(e) {
    e.preventDefault();
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    submitButtonLogin(email, password);
  });
}

function requestReset(e) {
  e.preventDefault();
  const email = document.getElementById("reset-email").value;
  fetch("http://localhost:8080/reset-password", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email })
  })
  .then(res => res.json())
  .then(data => {
    console.log("Resposta (Request Reset):", data);
    document.getElementById("step1").style.display = "none";
    document.getElementById("step2").style.display = "block";
  })
  .catch(err => console.error("Erro ao solicitar reset:", err));
}

function confirmReset(e) {
  e.preventDefault();
  const email = document.getElementById("reset-email-confirm").value;
  const token = document.getElementById("reset-pin").value;
  const newPassword = document.getElementById("reset-new-password").value;

  fetch("http://localhost:8080/validate-token", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      email,
      token,
      new_password: newPassword
    })
  })
  .then(res => res.json())
  .then(data => {
    console.log("Resposta (Confirm Reset):", data);
    alert("Senha alterada com sucesso!");
    window.location.href = "index.html";
  })
  .catch(err => console.error("Erro ao confirmar reset:", err));
}