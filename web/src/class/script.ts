export class script {
    id: number;
    username: string;
    name: string;
    type: string;
    language: string;
    code: string;
    cycle: number;
    status: boolean;
    create_time: number;
    update_time: number;
    description: string;
  
    constructor(
        id:number,username: string, name: string, type: string,
        language: string, code: string, cycle: number,
        status: boolean, create_time: number, update_time: number,
        description: string
    ) {
        this.id = 0;
        this.username = username;
        this.name = name;
        this.type = type;
        this.language = language;
        this.code = code;
        this.cycle = cycle;
        this.status = status;
        this.create_time = create_time;
        this.update_time = update_time;
        this.description = description;
    }
}