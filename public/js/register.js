const homePage = 'http://localhost:5050/'

let FormS = document.querySelector(".form-data");

FormS.addEventListener('submit', function (e) {

    e.preventDefault();

    let elem = e.target

    let formData = {
        familyName: elem.querySelector('[name="inputFamilyName"]').value,
        givenName: elem.querySelector('[name="inputGivenName"]').value,
        patronymic: elem.querySelector('[name="inputPatronymic"]').value,
        phone: elem.querySelector('[name="inputPhone"]').value,
        email: elem.querySelector('[name="inputEmail"]').value,
        password1: elem.querySelector('[name="password_1"]').value,
        password2: elem.querySelector('[name="password_2"]').value,
    }

    axios.post('http://localhost:5050/api/auth/register', {
        familyName: formData.familyName,
        givenName: formData.givenName,
        patronymic: formData.patronymic,
        phone: formData.phone,
        email: formData.email,
        password1: formData.password1,
        password2: formData.password2,

    })
        .then(function (response) {
            
            if (response.data.code === 200) {
                window.location.replace(homePage)
            }
        })

});
