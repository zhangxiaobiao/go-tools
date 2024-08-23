// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {check} from '../models';
import {fileinfo} from '../models';
import {utils} from '../models';

export function AllChecks(arg1:string):Promise<Array<check.Check>>;

export function BackUpFile(arg1:string):Promise<void>;

export function ChangeFileInfo(arg1:fileinfo.FileInfo):Promise<utils.Resp[interface {}]>;

export function DelChecks(arg1:string):Promise<number>;

export function DesktopPath():Promise<string>;

export function DeviceInfo():Promise<utils.Resp[interface {}]>;

export function EditHostsFile(arg1:string):Promise<utils.Resp[interface {}]>;

export function ImportBackup(arg1:string):Promise<utils.Resp[interface {}]>;

export function OpenFile():Promise<utils.Resp[interface {}]>;

export function OpenThisFile(arg1:string):Promise<utils.Resp[interface {}]>;

export function ReadHostsFile():Promise<utils.Resp[interface {}]>;

export function SaveFile(arg1:string,arg2:string):Promise<utils.Resp[interface {}]>;

export function SaveFileDialog(arg1:string):Promise<utils.Resp[interface {}]>;

export function SaveTime(arg1:string,arg2:string):Promise<number>;

export function SelectFile():Promise<utils.Resp[interface {}]>;
