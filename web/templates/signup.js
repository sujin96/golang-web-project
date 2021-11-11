
const idText = document.getElementById("id");
const nameText = document.getElementById("name");
const pwText = document.getElementById("pw");
const checkText = document.getElementById("check");


// const dayText = document.getElementById("day");
// const totaltimeText = document.getElementById("totaltime");
// const trytimeText = document.getElementById("trytime");
// const recoverytimeText = document.getElementById("recoverytime");
// const frontcountText = document.getElementById("frontcount");
// const backcountText = document.getElementById("backcount");
// const avgrpmText = document.getElementById("avgrpm");
// const avgspeedText = document.getElementById("avgspeed");
// const distanceText = document.getElementById("distance");
// const musclenumText = document.getElementById("musclenum");
// const KcalorynumText = document.getElementById("Kcalorynum");
// const areaText = document.getElementById("area");
// const birthText = document.getElementById("birth");
// const bike_infoText = document.getElementById("bike_info");
// const careerText = document.getElementById("career");
// const clubText = document.getElementById("club");
// const emailText = document.getElementById("email");






const submitBtn = document.getElementById("submitBtn");




const regex = / /gi;

submitBtn.disabled = true;

function checkBtnEnable(event){
    const targetId = event.target.id;

    switch(targetId) {
        case 'id':
        case 'name':
            if(event.target.value.replace(regex, "") == ""){
                event.target.style.backgroundColor = "red";
            }
            else{
                event.target.style.backgroundColor = "green";
            }
            if((idText.style.backgroundColor == "green") && (nameText.style.backgroundColor == "green") && (checkText.style.backgroundColor == "green")){
                submitBtn.disabled = false;
            }
            else{
                submitBtn.disabled = true;
            }
            break;
        case 'pw':
        case 'check':
            if(checkText.value != pwText.value){
                checkText.style.backgroundColor = "red";
            }
            else{
                checkText.style.backgroundColor = "green";
            }
            if((idText.style.backgroundColor == "green") && (nameText.style.backgroundColor == "green") && (checkText.style.backgroundColor == "green")){
                submitBtn.disabled = false;
            }
            else{
                submitBtn.disabled = true;
            }
            break;
            
        default:
            console.log("default");
            break;
    }
}

function init() {
    idText.addEventListener("input", checkBtnEnable);
    nameText.addEventListener("input", checkBtnEnable);
    pwText.addEventListener("input", checkBtnEnable);
    checkText.addEventListener("input", checkBtnEnable);

    // dayText.addEventListener("day", checkBtnEnable);
    // totaltimeText.addEventListener("totaltime", checkBtnEnable);
    // trytimeText.addEventListener("trytime", checkBtnEnable);
    // recoverytimeText.addEventListener("recoverytime", checkBtnEnable);
    // frontcountText.addEventListener("frontcount", checkBtnEnable);
    // backcountText.addEventListener("backcount", checkBtnEnable);
    // avgrpmText.addEventListener("avgrpm", checkBtnEnable);
    // avgspeedText.addEventListener("avgspeed", checkBtnEnable);
    // distanceText.addEventListener("distance", checkBtnEnable);
    // musclenumText.addEventListener("musclenum", checkBtnEnable);
    // KcalorynumText.addEventListener("Kcalorynum", checkBtnEnable);
    // areaText.addEventListener("area", checkBtnEnable);
    // birthText.addEventListener("birth", checkBtnEnable);
    // bike_infoText.addEventListener("bike_info", checkBtnEnable);
    // careerText.addEventListener("career", checkBtnEnable);
    // clubText.addEventListener("club", checkBtnEnable);
    // emailText.addEventListener("email", checkBtnEnable);


}

init();
