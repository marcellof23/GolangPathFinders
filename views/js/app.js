async function tangkapMarcello() {
    const url = '/api/wakwaw';
    try {
        let res = await fetch(url);
        console.log(res);
        console.log(await res.json())
        // return await res.json();
    } catch (error) {
        console.log(error);
    }
}

const marcello = document.getElementById("marcello");

marcello.addEventListener("click", e =>{
    tangkapMarcello();
})