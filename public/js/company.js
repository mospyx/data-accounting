async function getCompanies() {
    const response = await fetch("http://localhost:5050/api/companies", {
        headers: new Headers({
            'Authorization': localStorage.getItem("jwt")
        })
    });
    const data = await response.json();

    let arrID = new Array();
    for (let i = 0; i < data.length; i++) {
        const {ID} = data[i]
        arrID[i] = ID
    }
    initElements(arrID)
    for (let i = 0; i < data.length; i++  ) {
        const {ID} = data[i]
        await getCompany(ID)
    }

}

async function getCompany(ID) {
    const response = await fetch(`http://localhost:5050/api/companies/${ID}`, {
        headers: new Headers({
            'Authorization': localStorage.getItem("jwt")
        })
    });
    const data = await response.json();
    const {company_profile} = data;
    const {
        name,
        website,
        email,
        phone,
    } = company_profile

    document.getElementById(`company_name_${ID}`).textContent = name;
    document.getElementById(`company_website_${ID}`).textContent = website;
    document.getElementById(`company_email_${ID}`).textContent = email;
    document.getElementById(`company_phone_${ID}`).textContent = phone;
}

function initElements(arrID) {
    let row = document.createElement('div')
    row.setAttribute('class', 'row gx-4 gx-lg-5');
    let content = document.getElementById('content')
    row = content.appendChild(row);

    let arrCard = [];

    while (arrCard.length < arrID.length) {

        let col = document.createElement('div')
        col.setAttribute('class', 'col-md-4 mb-5');
        col = row.appendChild(col);

        let card = document.createElement('div')
        card.setAttribute('class', 'card h-100');
        card = col.appendChild(card);

        let cardHeader = document.createElement('div')
        cardHeader.setAttribute('class', 'card-header');
        cardHeader = card.appendChild(cardHeader);

        let cardTitle = document.createElement('h2')
        cardTitle.setAttribute('class', 'card-title');
        cardTitle = cardHeader.appendChild(cardTitle);

        //name

        let companyName = document.createElement('div')
        companyName.setAttribute('id', `company_name_${arrID[arrCard.length]}`);
        cardTitle.appendChild(companyName);

        let cardBody = document.createElement('div')
        cardBody.setAttribute('class', 'card-body');
        cardBody = card.appendChild(cardBody);

        //website

        let inputGroup = document.createElement('div')
        inputGroup.setAttribute('class', 'input-group');
        inputGroup = cardBody.appendChild(inputGroup);

        let inputGroupPrepend = document.createElement('div')
        inputGroupPrepend.setAttribute('class', 'input-group-prepend');
        inputGroupPrepend = inputGroup.appendChild(inputGroupPrepend);

        let inputGroupText = document.createElement('span')
        inputGroupText.setAttribute('class', 'input-group-text');
        inputGroupText = inputGroupPrepend.appendChild(inputGroupText)
        let groupText = document.createTextNode('🌐')
        inputGroupText.appendChild(groupText);

        let companyWebsite = document.createElement('span')
        companyWebsite.setAttribute('class', 'form-control');
        companyWebsite.setAttribute('id', `company_website_${arrID[arrCard.length]}`);
        inputGroup.appendChild(companyWebsite);

        //phone

        let inputGroup2 = document.createElement('div')
        inputGroup2.setAttribute('class', 'input-group');
        inputGroup2 = cardBody.appendChild(inputGroup2)

        let inputGroupPrepend2 = document.createElement('div')
        inputGroupPrepend2.setAttribute('class', 'input-group-prepend');
        inputGroupPrepend2 = inputGroup2.appendChild(inputGroupPrepend2)

        let inputGroupText2 = document.createElement('span')
        inputGroupText2.setAttribute('class', 'input-group-text');
        inputGroupText2 = inputGroupPrepend2.appendChild(inputGroupText2)
        let groupText2 = document.createTextNode('📞')
        inputGroupText2.appendChild(groupText2);

        let companyPhone = document.createElement('span')
        companyPhone.setAttribute('class', 'form-control');
        companyPhone.setAttribute('id', `company_phone_${arrID[arrCard.length]}`);
        inputGroup2.appendChild(companyPhone);

        //email

        let inputGroup3 = document.createElement('div')
        inputGroup3.setAttribute('class', 'input-group');
        inputGroup3 = cardBody.appendChild(inputGroup3)

        let inputGroupPrepend3 = document.createElement('div')
        inputGroupPrepend3.setAttribute('class', 'input-group-prepend');
        inputGroupPrepend3 = inputGroup3.appendChild(inputGroupPrepend3)

        let inputGroupText3 = document.createElement('span')
        inputGroupText3.setAttribute('class', 'input-group-text');
        inputGroupText3 = inputGroupPrepend3.appendChild(inputGroupText3)
        let groupText3 = document.createTextNode('✉')
        inputGroupText3.appendChild(groupText3);

        let companyEmail = document.createElement('span')
        companyEmail.setAttribute('class', 'form-control');
        companyEmail.setAttribute('id', `company_email_${arrID[arrCard.length]}`);
        inputGroup3.appendChild(companyEmail);

        let cardFooter = document.createElement('div')
        cardFooter.setAttribute('class', 'card-footer');
        cardFooter = card.appendChild(cardFooter)

        let buttenInfo = document.createElement('a')
        buttenInfo.setAttribute('class', 'btn btn-primary btn-sm');
        buttenInfo.setAttribute('href', '#!');
        buttenInfo = cardFooter.appendChild(buttenInfo)
        let buttenInfoText = document.createTextNode('More Info')
        buttenInfo.appendChild(buttenInfoText);

        let buttenEmployees = document.createElement('a')
        buttenEmployees.setAttribute('class', 'btn btn-primary btn-sm');
        buttenEmployees.setAttribute('href', '#!');
        buttenEmployees = cardFooter.appendChild(buttenEmployees)
        let buttenEmployeesText = document.createTextNode('Employees')
        buttenEmployees.appendChild(buttenEmployeesText);

        let buttenEdit = document.createElement('a')
        buttenEdit.setAttribute('class', 'btn btn-primary btn-sm');
        buttenEdit.setAttribute('href', '#!');
        buttenEdit = cardFooter.appendChild(buttenEdit)
        let buttenEditText = document.createTextNode('Edit')
        buttenEdit.appendChild(buttenEditText);

        arrCard.push(1)
    }
}

getCompanies()
