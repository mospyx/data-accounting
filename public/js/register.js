const homePage = 'http://localhost:5050/'

let FormS = document.querySelector(".form-data");

FormS.addEventListener('submit', function (e) {

    e.preventDefault();

    let elem = e.target

    let formData = {
        family_name: elem.querySelector('[name="family_name"]').value,
        given_name: elem.querySelector('[name="given_name"]').value,
        patronymic: elem.querySelector('[name="patronymic"]').value,
        phone: elem.querySelector('[name="phone"]').value,
        email: elem.querySelector('[name="email"]').value,
        password_1: elem.querySelector('[name="password_1"]').value,
        password_2: elem.querySelector('[name="password_2"]').value,
    }

    axios.post('http://localhost:5050/api/auth/register', {
        family_name: formData.family_name,
        given_name: formData.given_name,
        patronymic: formData.patronymic,
        phone: formData.phone,
        email: formData.email,
        password_1: formData.password_1,
        password_2: formData.password_2,

    })
        .then(function (response) {
            //TODO: Информировать о создании юзера и добавить действие
            if (response.data.id != 0) {
                window.location.replace(homePage)
            }
        })

});
