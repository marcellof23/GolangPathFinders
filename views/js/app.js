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

const graphForm = document.getElementById("graphForm");
const inpFile = document.getElementById("inpFile");

graphForm.addEventListener("submit", e =>{
    e.preventDefault();
    const endpoint = '/api/graph';
    const formData = new FormData();
    formData.append("inpFile",inpFile.files[0]);

    const reader = new FileReader();
    reader.onload = () => {
        const lines = reader.result.split('\n').map((line)=>{
            return line.split(' ');
        });
        console.log(lines);
    }
    reader.readAsText(inpFile.files[0]);

    console.log(inpFile.files[0]);
    (async function(){
        const res = await fetch(endpoint,{
            method: 'POST',
            body: formData,
        }).catch((err)=>console.log(err));
        
        console.log(await res.json());
    })()
});