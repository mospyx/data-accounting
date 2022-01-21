const homePage = 'http://localhost:5000/'

let FormS = document.querySelector(".form-data");

FormS.addEventListener('submit', function (e) {

    e.preventDefault();

    let elem = e.target

    let formData = {
        email: elem.querySelector('[name="email"]').value,
        password: elem.querySelector('[name="password"]').value,
    }

    axios.post('http://localhost:5000/api/auth/login', {
        email: formData.email,
        password: formData.password,
    })
        .then(function (response) {
            let token = response.data.token;
            localStorage.setItem("jwt", 'Bearer ' + token);
            axios.defaults.headers.common['Authorization'] = 'Bearer ' + token;
            if (response.data.code === 200) {
                window.location.replace(homePage)
            }
        })

});
