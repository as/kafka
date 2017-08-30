// This file contains the protocol definitions for the latest version of the Kafka protocol, found here: 
//    https://cwiki.apache.org/confluence/display/KAFKA/A+Guide+To+The+Kafka+Protocol#AGuideToTheKafkaProtocol-Partitioningandbootstrapping
//
// In order to generate Go from this file, you need wire9/cmd (rename the executable to wire9)
//   github.com/as/wire9/cmd
// wire9 -f kafka_wire9.go .
//
// To insert print statements into the boilerplate, run wire9 with -d
//
package kafka

type varint int

//wire9 Str N[4] data[N]

//wire9 TX  ApiKey[2] ApiVer[2] CorrelationID[4] ClientID[,Str]
//wire9 MetaDataTX     TX[,TX] 
//wire9 ProduceTX      TX[,TX] 
//wire9 FetchTX        TX[,TX] 
//wire9 OffsetTX       TX[,TX] 
//wire9 OffsetCommitTX TX[,TX] 
//wire9 OffsetFetchTX  TX[,TX]

//wire9 RX CorrelationID[4] ClientID[,Str]
//wire9 MetaDataRX     RX[,RX] 
//wire9 ProduceRX      RX[,RX] 
//wire9 FetchRX        RX[,RX] 
//wire9 OffsetRX       RX[,RX] 
//wire9 OffsetCommitRX RX[,RX] 
//wire9 OffsetFetchRX  RX[,RX]

//!wire9 Message    CRC[int32] Magic[1] Attr[1] Key[,Str] Value[,Str]
//!wire9 MessageSet Offset[int64] Size[int32]

//wire9 Tag         KeyLen[,varint] Key[,Str] ValueLen[,varint] Value[,Str]
//wire9 Headers     N[4] Header[N,[]Tag]
//wire9 Record      Len[,varint] Attr[1] TimeDelta[,varint] OffsetDelta[,varint] Tag[,Tag] Headers[,Headers]
//wire9 RecordBatch Offset[8] Len[8] LeaderEpoch[4] Magic[1] Attr[2] LastOffsetDelta[4] FirstTime[8] MaxTime[8] ProdID[8] ProdEpoch[2] FirstSeq[4] N[4] Records[N,[]Record]
//wire9 Broker      ID[4] Host[,Str] Port[4]
//wire9 MetaPart    Err[2] ID[4] Leader[4] Replicas[4] ISR[4]
//wire9 MetaTopic   Err[2] Topic[,Str] N[4] MetaParts[N,[]MetaPart]

//wire9 MetaTopicTX   N[4] Topics[N,[]Str]
//wire9 MetaTopicRX   N[4] Brokers[N,[]Broker] M[4] Topics[M, []MetaTopic]

// Producer API

//wire9 ProdTXPart    ID[4] MessageSetSize[4] MessageSet[4]
//wire9 ProdTXTopic   Name[,Str]         N[4]  Parts[N,[]ProdTxPart]
//wire9 ProdTX        NeedAcks[2] Timeout[4] N[4] Topics[N,[]ProdTXTopic] 

//wire9 ProdRXPart    ID[4] Err[2] Offset[8] Time[8]
//wire9 ProdRXTopic   Name[,Str] N[4] Parts[N,[]ProdRXPart]
//wire9 ProdRX        N[4] Topics[N,[]ProdRXTopic] ThrottleTime[4] 

// Fetch API

//wire9 FetchTXPart   ID[4] Offset[8] MaxBytes[4]
//wire9 FetchTXTopic  Name[,Str] N[4] Parts[N,[]FetchTXPart]
//wire9 FetchTX       ReplicaId[4] MaxWait[4] MinBytes[4] N[4] Topic[N,[]FetchTXTopic]

//wire9 FetchRXPart   ID[4] Err[2] HiWater[8] MessageSetSize[4] MessageSet[4]
//wire9 FetchRXTopic  Name[,Str] N[4] Parts[N,[]FetchRXPart]
//wire9 FetchRX       ThrottleTime[4] N[4] Topics[N,[]FetchRXTopic] 

// Offset API

//wire9 OffsetTXPart ID[4] Time[8] MaxOffsets[4]
//wire9 OffsetTXTopic Name[,Str] N[4] Parts[N,[]OffsetTXPart]
//wire9 OffsetTX      ReplicaId[4] N[4] Topics[N,[]OffsetTXTopic]

//wire9 Offset        Q0[8]
//wire9 OffsetRXPart   ID[4] Err[2] Time[8] N[4] Offsets[N,[]Offset]
//wire9 OffsetRXTopic  Name[,Str] N[4] Parts[N, []OffsetRXPart]
//wire9 OffsetRX       N[4] Topics[N,[]OffsetRXTopic]

// Offset Commit API

//wire9 GroupCoTX   Group[,Str]
//wire9 GroupCoRX   Err[2] ID[4] Host[,Str] Port[4]

//wire9 CommitTXPart  ID[4] Offset[8] MetaData[,Str]
//wire9 CommitTXTopic Name[,Str] N[4] Parts[N,[]CommitTXPart]
//wire9 CommitTX      Group[,Str] Gen[4] Cons[,Str] Retain[8] N[4] Topics[N,[]CommitTXTopic] 

//wire9 CommitRXPart   ID[4] Err[2]
//wire9 CommitRXTopic  Name[,Str] N[4] Parts[N,[]CommitRXPart]
//wire9 CommitRX       N[4] Topic[N,CommitRXTopic]

// Offset Fetch API

//wire9 FetchTXPart  ID[4]
//wire9 FetchTXTopic Name[,Str]  N[4] Parition[N,[]FetchTXPart]
//wire9 FetchTX      Group[,Str] N[4] Topic[N,[]FetchTXTopic]

//wire9 FetchRxPart  ID[4] Offset[8] MetaData[,Str] Err[2]
//wire9 FetchRxTopic Name[,Str] N[4] Parts[N,[]FetchTXPart]
//wire9 FetchRX      N[4] Topic[N,FetchRxTopic]

// Group Membership API

//wire9  Proto Name[,Str] MetaData[,Str]
//wire9  DefaultProto Ver[2] N[4] TopicNames[N,[]Str]

//wire9  JoinGroupTX  Group[,Str] Timeout[4] RebalanceTimeout[4]    Member[,Str] Prototype[,Str] N[4]  Protos[N,[]Proto]
//wire9  JoinGroupRX  Err[2] GenID[4] Proto[,Str] Leader[,Str]      Member[,Str]                 N[4] Members[N,[]Proto]

//wire9  AssignPart   ID[4]
//wire9  AssignTopic  N[4] Parts[N,[]AssignPart]
//wire9  AssignMember Ver[2]                          N[4] Parts[N,   []AssignTopic]  UserData[,Str]
//wire9  GroupAssign  Member[,Str]                    N[4] Assigns[N, []AssignMember]
//wire9  SyncGroupTX  Group[,Str] Gen[4] Member[,Str] N[4] Assigns[N, []GroupAssign]
//wire9  SyncGroupRX  Err[2]                          N[4] Assigns[N, []AssignMember]

//wire9 LeaveGroupTX Group[,Str] Member[,Str]
//wire9 LeaveGroupRX Err[2]

// Heartbeat API

//wire9 HeartTX Group[,Str] Gen[4] Member[,Str]
//wire9 HeartRX Err[2]

// Admin API

//wire9 GroupList Group[,Str] Prototype[,Str]
//wire9 ListGroupsTX Zero[1]
//wire9 ListGroupsRX Err[2] N[4] Groups[N,[]GroupList]

//wire9 DescGroupsTX N[4] Groups[N,[]Str]
//wire9 DescGroupMember Name[,Str] Client[,Str] Host[,Str] MetaData[,Str] N[4] Assigns[N, []GroupAssign]
//wire9 DescGroup       Err[2] Group[,Str] State[,Str] Prototype[,Str] Proto[,Str] N[4] Members[N, []DescGroupMember]
//wire9 DescGroupsRX    Err[2] Group[,Str] 
