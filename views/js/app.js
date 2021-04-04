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

tangkapMarcello();

const graphForm = document.getElementById("graphForm");

graphForm.addEventListener("submit", e =>{
    e.preventDefault();
    console.log(e);
});