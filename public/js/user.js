// axios('http://localhost:5000/api/user/current_user', {
//     headers: new Headers({
//         'Authorization': localStorage.getItem("jwt")
//     })
// })
//     .then(function (response) {
//         document.getElementById('currentUser').textContent = JSON.stringify(response.data.email)
//     })

// fetch("http://localhost:5000/api/user/current_user", {
//     headers: new Headers({
//         'Authorization': localStorage.getItem("jwt")
//     })
//     })
//         .then(function (response) {
//             const data = response.json();
//             const {email} = data;
//             document.getElementById("current_user").textContent = email
//         })


async function getCurrentUser() {
    const response = await fetch("http://localhost:5000/api/user/current_user", {
        headers: new Headers({
            'Authorization': localStorage.getItem("jwt")
        })
    });
    const data = await response.json();
    const {email} = data;

    document.getElementById("current_user").textContent = email;
}

getCurrentUser()
