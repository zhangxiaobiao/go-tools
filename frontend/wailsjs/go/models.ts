export namespace check {
	
	export class Check {
	    id: number;
	    date: string;
	    start: string;
	    end: string;
	
	    static createFrom(source: any = {}) {
	        return new Check(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.date = source["date"];
	        this.start = source["start"];
	        this.end = source["end"];
	    }
	}

}

export namespace fileinfo {
	
	export class FileInfo {
	    filePath: string;
	    createTime: number;
	    accessTime: number;
	    updateTime: number;
	
	    static createFrom(source: any = {}) {
	        return new FileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.filePath = source["filePath"];
	        this.createTime = source["createTime"];
	        this.accessTime = source["accessTime"];
	        this.updateTime = source["updateTime"];
	    }
	}

}

export namespace utils {
	
	export class Resp[interface {}] {
	    code: number;
	    msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Resp[interface {}](source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}

}

