var dropzone;
var target;
var uploadInput;

var fileCtxtDOM = "<a download target='_blank' id='ctxt-download'>⇣ Download</a><a>📂 Preview / Edit</a><div></div><a id='ctxt-delete'>✖ Delete</a><a>✎ Rename</a><div></div><a>� Info</a><a>🔗 Get Link</a>";
var dropCtxtDOM = "<a onclick='selectFile();'>⇡ Upload</a><a>✚ New Text File</a>";

function startup() {
    ctxtModal = new Modal("context-menu");
    errModal = new Modal("error");
    dlogModal = new Modal("dialog").setAnim("fadedropin", "fadedropout").set("<input placeholder='rename things' class='vezinput'><button class='pi-button' style='background: aqua;'>click</button>").show();
    dlogBack = new Modal("dialog-back").setAnim("fadein").show();
    dlogBack.dom.onclick = function() {
        dlogModal.hide();
        dlogBack.hide();
    }
    
    dropzone  = document.getElementById('dropzone');
    uploadInput = document.getElementById("uploadfile");
    
    //context menu stuff
    document.addEventListener('contextmenu', function(e) {
        if (e.target.className.split(" ")[0] == "file") {
            target = e.target;
            filename = target.parentNode.textContent+target.getAttribute("filetype");

            ctxtModal.dom.innerHTML = fileCtxtDOM;
            ctxtModal.dom.querySelector("#ctxt-download").href = "/dnl/" + filename;
            ctxtModal.dom.querySelector("#ctxt-delete").onclick = function() { deleteFile(filename) };
            console.log("deleteFile('" + filename + "')");
            
            var view_bottom = window.innerHeight + window.pageYOffset; //bottom of viewport
            if (view_bottom - e.pageY > 300) {
                ctxtModal.dom.style.top = e.pageY+"px";
                ctxtModal.dom.style.left = e.pageX+"px";
            } else {
                ctxtModal.dom.style.top = (e.pageY - 300)+"px";
                ctxtModal.dom.style.left = e.pageX+"px";
            }
            ctxtModal.show();
            e.preventDefault();
            
        } else if (e.target.id == "dropzone") {
            ctxtModal.dom.innerHTML = dropCtxtDOM;
            
            var view_bottom = window.innerHeight + window.pageYOffset;
            if (view_bottom - e.pageY < 100) {
                ctxtModal.dom.style.top = (e.pageY-100)+"px";
                ctxtModal.dom.style.left = e.pageX+"px";
            } else {
                ctxtModal.dom.style.top = e.pageY+"px";
                ctxtModal.dom.style.left = e.pageX+"px";
            }
            ctxtModal.show();
            e.preventDefault();
            
        }else {
            ctxtModal.hide();
            //e.preventDefault();
        }
    }, false);
    
    document.addEventListener('click', function(e) {
        ctxtModal.hide();
    });
    
    //file upload stuff
	dropzone.addEventListener('dragover', cancel, false);
	dropzone.addEventListener('dragenter', cancel, false);
	dropzone.addEventListener('drop', function(e) { 
		e = e || window.event;
		if (e.preventDefault) { e.preventDefault(); }

		var fs = e.dataTransfer.files;
		uploadFiles(fs)
		return false;
	});
	uploadInput.addEventListener('change', function(e) {
		uploadFiles(uploadInput.files);
		uploadInput.value = null;
	});
}

Element.prototype.remove = function() {
    this.parentElement.removeChild(this);
}
NodeList.prototype.remove = HTMLCollection.prototype.remove = function() {
    for(var i = this.length - 1; i >= 0; i--) {
        if(this[i] && this[i].parentElement) {
            this[i].parentElement.removeChild(this[i]);
        }
    }
}

function deleteFileElement(eldel) {
    eldel.parentNode.style.animationDirection = "reverse";
    eldel.parentNode.style.animationName = "expand";
    
    fun = function (elem) {
        return function(e) {
            elem.parentNode.remove();
            elem.remove();
        };
    }(eldel);
    eldel.parentNode.addEventListener("animationend", fun);
}

function addFileElement(name, ext) {
    dropzone.insertAdjacentHTML('afterbegin', "<a href='/upl/"+ name + "." +ext +"'><div class='file "+ ext +"' filetype='."+ext+"'></div>"+name+"</a>");
    dropzone.firstChild.style.animationName = "expand";

    fun = function (elem) {
        return function(e) {
            elem.style.animationName = "";
        };
    }(dropzone.firstChild);
    dropzone.firstChild.addEventListener("animationend", fun);
}

function selectFile() {
    uploadInput.click();
}

//file upload stuff

function cancel(e) {
	if (e.preventDefault) { 
		e.preventDefault(); 
	}
	return false;
}

function uploadFiles(fs) {
	if (fs.length == 0) {
		console.log('no file to upload!');
		return;
	}
	var fd = new FormData();
	var xhr = new XMLHttpRequest();
	    xhr.open('POST', '/', true);
	
    var fsinfo = [];
    for (var i = 0; i < fs.length; i++) { 
        fsinfo.push({name: fs[i].name.substr(0, fs[i].name.lastIndexOf('.')), ext: fs[i].name.substr(fs[i].name.lastIndexOf('.') + 1)});
    }
    
    xhr.onreadystatechange = function(files) {
        return function() {
            if (xhr.readyState == 4) {
                if (xhr.status == 0) {
                    console.log('A state error occurred!');
                } else if (xhr.status == 200) {
                    console.log('file uploaded! ');
                    for (var i = 0; i < files.length; i++) {
                        addFileElement(files[i].name, files[i].ext);
                    }
                }
            }
        };
    }(fsinfo);
	    
    for (var i = 0; i < fs.length; i++) {
        fd.append('uploadfile', fs[i], fs[i].name);
    }

	xhr.send(fd);
}

function deleteFile(name) {
	var xhr = new XMLHttpRequest();
	    xhr.open('DELETE', "/upl/" + name, true);
    //apply to target
    xhr.onreadystatechange =  function(eldel) {
        return function() {
                if (xhr.readyState == 4) {
                    if (xhr.status == 0) {
                        console.log('A state error occurred!');
                    } else if (xhr.status == 200) {
                        console.log('file deleted!');
                        deleteFileElement(eldel);
                    }
                }
            };
    }(target);

    xhr.send();
}