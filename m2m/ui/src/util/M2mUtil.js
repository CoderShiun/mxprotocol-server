export const SUPER_ADMIN = '0';

export function BackToLora() {
    window.location.replace(getLoraHost());
}

export function getLoraHost(){
    let host = process.env.REACT_APP_LORA_APP_SERVER;
    const origin = window.location.origin;
    
    if(origin.includes(process.env.REACT_APP_SUBDOM_M2M)){
        host = origin.replace(process.env.REACT_APP_SUBDOM_M2M, process.env.REACT_APP_SUBDOM_LORA);
    }

    return host;
}