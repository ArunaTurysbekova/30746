var images = ["main/175706_or.jpg", "main/993008.jpg", "main/Unknown.jpeg"];

var myImagesDiv = document.getElementById("myImages");

function showImages() {
  myImagesDiv.innerHTML = "";
  for (var i = 0; i < images.length; i++) {
    var img = document.createElement("img");
    img.src = images[i];
    myImagesDiv.appendChild(img);
  }
}

function exp(b){
  return b;
}
console.log(exp(12))

function exp1(b){
  return b;
}
console.log(exp1({
  0:"text1",
  1:"text2",
  2:"text3",
  3:"text4",
}))

