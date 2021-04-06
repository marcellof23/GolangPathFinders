window.markers = [];
window.polylines = [];

let activeMarkers = [];
let pointedMarkers = [];

let sourceMarker = null;
let destinationMarker = null;

const lineSymbol = {
	path: "M 0,-1 0,1",
	strokeOpacity: 1,
	scale: 4,
};

function ClearControl(controlDiv) {
	const controlUI = document.createElement("div");
	controlUI.style.backgroundColor = "#dc3444";
	controlUI.style.border = "2px solid #dc3444";
	controlUI.style.borderRadius = "3px";
	controlUI.style.boxShadow = "0 2px 6px rgba(0,0,0,.3)";
	controlUI.style.cursor = "pointer";
	controlUI.style.marginTop = "8px";
	controlUI.style.marginBottom = "22px";
	controlUI.style.textAlign = "center";
	controlUI.title = "Click to remove markers and reset the map";
	controlDiv.appendChild(controlUI);

	const controlText = document.createElement("div");
	controlText.style.color = "#FFF";
	controlText.style.fontFamily = "Roboto,Arial,sans-serif";
	controlText.style.fontSize = "16px";
	controlText.style.lineHeight = "38px";
	controlText.style.paddingLeft = "5px";
	controlText.style.paddingRight = "5px";
	controlText.innerHTML = "Clear Nodes";
	controlUI.appendChild(controlText);

	controlUI.addEventListener("click", () => {
		clearMarker();
	});
}

function SubmitControl(controlDiv) {
	const controlUI = document.createElement("div");
	controlUI.style.backgroundColor = "#48ac44";
	controlUI.style.border = "2px solid #48ac44";
	controlUI.style.borderRadius = "3px";
	controlUI.style.boxShadow = "0 2px 6px rgba(0,0,0,.3)";
	controlUI.style.cursor = "pointer";
	controlUI.style.marginTop = "8px";
	controlUI.style.marginBottom = "22px";
	controlUI.style.textAlign = "center";
	controlUI.title = "Click to recenter the map";
	controlDiv.appendChild(controlUI);

	const controlText = document.createElement("div");
	controlText.style.color = "#FFF";
	controlText.style.fontFamily = "Roboto,Arial,sans-serif";
	controlText.style.fontSize = "16px";
	controlText.style.lineHeight = "38px";
	controlText.style.paddingLeft = "5px";
	controlText.style.paddingRight = "5px";
	controlText.innerHTML = "Find Shortest Path";
	controlUI.appendChild(controlText);

	controlUI.addEventListener("click", () => {
		handleFetchShortestPath();
	});
}

function initMap() {
	const bandung = { lat: -6.9175, lng: 107.6191 };
	google.maps.event.trigger(map, "resize");
	window.map = new google.maps.Map(document.getElementById("map"), {
		zoom: 17,
		center: bandung,
	});

	window.map.addListener("click", (e) => {
		placeMarkerAndPanTo({ lat: e.latLng.lat(), lng: e.latLng.lng() });
	});

	const clearControlDiv = document.createElement("div");
	const submitControlDiv = document.createElement("div");
	ClearControl(clearControlDiv);
	SubmitControl(submitControlDiv);
	window.map.controls[google.maps.ControlPosition.TOP_CENTER].push(
		clearControlDiv
	);
	window.map.controls[google.maps.ControlPosition.TOP_CENTER].push(
		submitControlDiv
	);
}

const graphForm = document.getElementById("graphForm");
const inpFile = document.getElementById("inpFile");

graphForm.onchange = () => {
	document.getElementById("inpFileLabel").innerHTML = inpFile.files[0].name;
};

graphForm.addEventListener("submit", (e) => {
	e.preventDefault();
	const reader = new FileReader();
	reader.onload = () => {
		clearMarker();

		const lines = reader.result.split("\n").map((line) => {
			return line.split(",");
		});
		for (let i = 1; i < 1 + parseInt(lines[0][0]); i++) {
			placeMarkerAndPanTo({
				lat: parseFloat(lines[i][1]),
				lng: parseFloat(lines[i][2]),
			});
		}
		for (
			let i = 1 + parseInt(lines[0][0]);
			i < 1 + 2 * parseInt(lines[0][0]);
			i++
		) {
			let row = lines[i][0].split(" ");
			for (let j = 0; j < parseInt(lines[0][0]); j++) {
				if (parseInt(row[j]) === parseInt("1")) {
					addLines(
						{
							lat: markers[i - 1 - parseInt(lines[0][0])].position.lat(),
							lng: markers[i - 1 - parseInt(lines[0][0])].position.lng(),
						},
						{ lat: markers[j].position.lat(), lng: markers[j].position.lng() }
					);
				}
			}
		}
	};
	reader.readAsText(inpFile.files[0]);
});

const placeMarkerAndPanTo = (coords) => {
	let marker = new google.maps.Marker({
		position: coords,
		map: window.map,
		icon: "http://maps.google.com/mapfiles/ms/icons/red-dot.png",
	});
	marker.addListener("click", (e) => {
		handleNodeLeftClick(marker);
	});
	marker.addListener("contextmenu", (e) => {
		handleNodeRightClick(marker);
	});
	marker.addListener("dblclick", (e) => {
		handleNodeDoubleClick(marker);
	});

	window.markers.push(marker);
	window.map.panTo(coords);
};

const clearMarker = () => {
	for (let i = 0; i < window.markers.length; i++) {
		window.markers[i].setMap(null);
	}
	for (let i = 0; i < window.polylines.length; i++) {
		window.polylines[i].setMap(null);
	}

	window.markers = [];
	window.polylines = [];
};

const addLines = (sourceCoords, destinationCoords) => {
	let polyline = new google.maps.Polyline({
		path: [sourceCoords, destinationCoords],
		geodesic: true,
		strokeOpacity: 0,
		icons: [
			{
				icon: lineSymbol,
				offset: "0",
				repeat: "30px",
			},
		],
	});
	polyline.addListener("contextmenu", (e) => {
		handlePolylineRightClick(polyline);
	});
	polyline.setMap(window.map);
	window.polylines.push(polyline);
};

const handleNodeLeftClick = (marker) => {
	marker.setIcon("http://maps.google.com/mapfiles/ms/icons/green-dot.png");
	if (pointedMarkers.length == 1) {
		addLines(
			{
				lat: pointedMarkers[0].position.lat(),
				lng: pointedMarkers[0].position.lng(),
			},
			{ lat: marker.position.lat(), lng: marker.position.lng() }
		);
		addLines(
			{ lat: marker.position.lat(), lng: marker.position.lng() },
			{
				lat: pointedMarkers[0].position.lat(),
				lng: pointedMarkers[0].position.lng(),
			},
		)
		marker.setIcon("http://maps.google.com/mapfiles/ms/icons/red-dot.png");
		pointedMarkers[0].setIcon(
			"http://maps.google.com/mapfiles/ms/icons/red-dot.png"
		);
		pointedMarkers = [];
	} else {
		pointedMarkers.push(marker);
	}
};

const handleNodeRightClick = (marker) => {
	if (activeMarkers.length == 2) {
		activeMarkers[0].setIcon(
			"http://maps.google.com/mapfiles/ms/icons/red-dot.png"
		);
		activeMarkers.shift();
	}
	activeMarkers.push(marker);
	sourceMarker = activeMarkers[0];
	destinationMarker = activeMarkers[1];
	marker.setIcon("http://maps.google.com/mapfiles/ms/icons/blue-dot.png");
};

const handleNodeDoubleClick = (marker) => {
	marker.setMap(null);
	window.markers = window.markers.filter((m) => {
		return (
			m.position.lat() != marker.position.lat() ||
			m.position.lng() != marker.position.lng()
		);
	});
};

const handlePolylineRightClick = (polyline) => {
	window.polylines = window.polylines.filter((p)=>{
		if(polyline.getPath().Nb[0].lat()==p.getPath().Nb[0].lat() && polyline.getPath().Nb[0].lng()==p.getPath().Nb[0].lng() && 
		polyline.getPath().Nb[1].lat()==p.getPath().Nb[1].lat() && polyline.getPath().Nb[1].lng()==p.getPath().Nb[1].lng() || (polyline.getPath().Nb[1].lat()==p.getPath().Nb[0].lat() && polyline.getPath().Nb[1].lng()==p.getPath().Nb[0].lng() && 
		polyline.getPath().Nb[0].lat()==p.getPath().Nb[1].lat() && polyline.getPath().Nb[0].lng()==p.getPath().Nb[1].lng())){
			p.setMap(null);
		}
		return !(polyline.getPath().Nb[0].lat()==p.getPath().Nb[0].lat() && polyline.getPath().Nb[0].lng()==p.getPath().Nb[0].lng() && 
		polyline.getPath().Nb[1].lat()==p.getPath().Nb[1].lat() && polyline.getPath().Nb[1].lng()==p.getPath().Nb[1].lng()) 
		&& 
		!(polyline.getPath().Nb[1].lat()==p.getPath().Nb[0].lat() && polyline.getPath().Nb[1].lng()==p.getPath().Nb[0].lng() && 
		polyline.getPath().Nb[0].lat()==p.getPath().Nb[1].lat() && polyline.getPath().Nb[0].lng()==p.getPath().Nb[1].lng())
	})
	// console.log(polyline.getPath());
	// console.table(polyline.getPath());
	// console.log(polyline.getPath().Nb[0].lat() , polyline.getPath().Nb[0].lng());
	// console.log(polyline.getPath().Nb[1].lat() , polyline.getPath().Nb[1].lng());
	console.log(window.polylines)
};

function isMarkerEqual(firstMarker, secondMarker){
	return firstMarker.position.lat()==secondMarker.position.lat() && firstMarker.position.lng()==secondMarker.position.lng()
}

const handleFetchShortestPath = async ()=>{
	if(!sourceMarker || !destinationMarker){
		alert("Please choose a source and destination node");
	} else{
		console.log(sourceMarker.position.lat(), sourceMarker.position.lng());
		console.log(destinationMarker.position.lat(), destinationMarker.position.lng());
		let sourceNode = -1;
		let destNode = -1;
		for(let i=0;i<markers.length;i++){
			if(isMarkerEqual(markers[i],activeMarkers[0])){
				sourceNode = i;
				break;
			}
		}
		for(let i=0;i<markers.length;i++){
			if(isMarkerEqual(markers[i],activeMarkers[1])){
				destNode = i;
				break;
			}
		}
		let boolMatrix = [...Array(markers.length)].map(x=>Array(markers.length).fill(0))
		console.table(boolMatrix);
		for(let i=0;i<markers.length;i++){
			for(let j=0;j<markers.length;j++){
				for(let k=0;k<polylines.length;k++){
					let polyline = polylines[k];
					if(polyline.getPath().Nb[0].lat()==markers[i].position.lat() && polyline.getPath().Nb[0].lng()==markers[i].position.lng() && 
					polyline.getPath().Nb[1].lat()==markers[j].position.lat() && polyline.getPath().Nb[1].lng()==markers[j].position.lng() || (polyline.getPath().Nb[1].lat()==markers[i].position.lat() && polyline.getPath().Nb[1].lng()==markers[i].position.lng() && 
					polyline.getPath().Nb[0].lat()==markers[j].position.lat() && polyline.getPath().Nb[0].lng()==markers[j].position.lng())){
						boolMatrix[i][j] = 1;
					}
				}
			}
		}
		let matrixRelation = "";
		for(let i=0;i<markers.length;i++){
			for(let j=0;j<markers.length;j++){
				if(j==markers.length-1){
					matrixRelation += `${boolMatrix[i][j]},`
				} else{
					matrixRelation += `${boolMatrix[i][j]} `
				}
			}
		}
		console.table(boolMatrix);
		let coordinates = "";
		for(let i=0;i<markers.length;i++){
			coordinates+= `${markers[i].position.lat()} ${markers[i].position.lng()},`
		}
		console.log({
			numnodes: markers.length,
			sourcenode: sourceNode,
			destnode: destNode,
			matrixrelation: matrixRelation.slice(0,-1),
			coordinates: coordinates.slice(0,-1)
		});
		const res = await postData("http://localhost:5000/api/graphdata",{
			numnodes: markers.length,
			sourcenode: sourceNode,
			destnode: destNode,
			matrixrelation: matrixRelation.slice(0,-1),
			coordinates: coordinates.slice(0,-1)
		})
		console.log(res);
		let indexes = res.Path.split(",");
		console.log(indexes);
		for(let i=0;i<indexes.length;i++){
			window.markers[parseInt(indexes[i])].setIcon("http://maps.google.com/mapfiles/ms/icons/green-dot.png");
		}
	}
}

async function postData(url = '', data = {}) {
  // Default options are marked with *
  const response = await fetch(url, {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json'
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    body: JSON.stringify(data)
  });
  return response.json();
}
