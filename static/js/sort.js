const toSort = (toSortingElement) => {

    let category = toSortingElement.getAttribute('category');
    let elementAttribute = toSortingElement.getAttribute('element');
    let isAlreadyClicked = toSortingElement.getAttribute('is-already-clicked');

    if (category == null) {
        if (isAlreadyClicked == "false") {
            allJson.sort((a,b) => {
                return (a[elementAttribute] > b[elementAttribute])?1:-1;
            });
    
            toSortingElement.setAttribute('is-already-clicked', "true");
            
        } else {
            allJson.sort((a,b) => {
                return (a[elementAttribute] < b[elementAttribute])?1:-1;
            });
    
            toSortingElement.setAttribute('is-already-clicked', "false");
            
        }

    } else {
        if (isAlreadyClicked == "false") {

            allJson.sort((a,b) => {

                if (Array.isArray(a[category][elementAttribute])) {
                    return a[category][elementAttribute][0].localeCompare(b[category][elementAttribute][0], undefined, {numeric: true});
                }

                if (a[category][elementAttribute] == "" || a[category][elementAttribute] == "-") {
                    return 1;
                } else if (b[category][elementAttribute] == "" || b[category][elementAttribute] == "-") {
                    return -1;
                }

                return (a[category][elementAttribute] > b[category][elementAttribute])?1:-1;
            });
    
            toSortingElement.setAttribute('is-already-clicked', "true");
            
        } else {

            allJson.sort((a,b) => {

                if (Array.isArray(a[category][elementAttribute])) {
                    return b[category][elementAttribute][0].localeCompare(a[category][elementAttribute][0], undefined, {numeric: true});
                }

                if (a[category][elementAttribute] == "" || a[category][elementAttribute] == "-") {
                    return 1;
                } else if (b[category][elementAttribute] == "" || b[category][elementAttribute] == "-") {
                    return -1;
                }

                return (a[category][elementAttribute] < b[category][elementAttribute])?1:-1;
            });
    
            toSortingElement.setAttribute('is-already-clicked', "false");
            
        }
    }


    for (th of allElementOfHeadBar) {
        if(th != toSortingElement) {
            th.setAttribute('is-already-clicked', "false");
        }      
    }

    display();
}
