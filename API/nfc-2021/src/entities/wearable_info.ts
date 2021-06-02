import { Entity, Column, PrimaryGeneratedColumn, Generated } from 'typeorm';
import {resolveSrv} from "dns";

@Entity()
export class wearable_info {
  @PrimaryGeneratedColumn('uuid')
  wearable_SN: string;

  @Column()
  UUID: string;

  @Column()
  email: string;
  @Column()
  displayname: string;

  @Column()
  token: string;

  @Column()
  PSN: string;

  @Column()
  PSN_img: string;

  @Column()
  Ls_admin: string;

  @Column()
  weearable_SN: string;
}
