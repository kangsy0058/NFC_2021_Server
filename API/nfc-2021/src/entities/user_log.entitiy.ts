import { Entity, Column, PrimaryGeneratedColumn } from 'typeorm';

@Entity()
export class user_log {
  @PrimaryGeneratedColumn()
  IDX: number;

  @Column()
  wearable_SN: string;

  @Column()
  kiosk_SN: string;

  @Column('datetime')
  time: Date;

  @Column('datetime')
  date: Date;

  @Column()
  temp: string;

  @Column()
  detail_position: string;

  @Column()
  building_name: string;

  @Column()
  latitude: number;

  @Column()
  longitude: number;
}
