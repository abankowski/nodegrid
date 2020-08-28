package node

import (
	"testing"
	"encoding/json"
	)


func Test(t *testing.T) {

	jsonBytes := []byte("{\"metrics\":{\"transactionService_unknown_size\":\"64\",\"genesisAccepted\":\"true\",\"cluster_ownJoinedHeight\":\"819498\",\"nodeState_success\":\"69\",\"totalNumCBsInShapshots\":\"739538\",\"lastSnapshotHash\":\"2ccd7afb291e0fb9ec31dc098a1b60744cfd36586e6751fedc9be6754f9b6eb3\",\"nodeState\":\"Ready\",\"checkpointTipsRemoved\":\"589978\",\"redownload_maxCreatedSnapshotHeight\":\"829736\",\"peerApiRXFinishedCheckpoint\":\"1923147\",\"snapshotAttempt_failure\":\"4\",\"consensus_participateInRound\":\"34135\",\"nodeStartTimeMS\":\"1598502472879\",\"transactionAccepted\":\"2757798\",\"externalHost\":\"161.35.17.54\",\"snapshotAttempt_heightIntervalNotMet\":\"5708\",\"redownload_lastMajorityStateHeight\":\"829708\",\"acceptMajorityCheckpointBlockUniquesCount_1\":\"45830\",\"resolveMajorityCheckpointBlockUniquesCount_1\":\"45830\",\"TPS_all\":\"23.232900626890537\",\"resolveMajorityCheckpointBlockProposalCount_3\":\"45830\",\"parentSOEServiceQueryFailed\":\"113\",\"peerAddedFromRegistrationFlow\":\"69\",\"rewards_snapshotReward\":\"32921810681\",\"rewards_snapshotRewardWithoutStardust\":\"29629629612\",\"trustDataPollingRound\":\"1831\",\"blacklistedAddressesSize\":\"0\",\"rewards_stardustBalanceAfterReward\":\"1025540740340764\",\"consensus_participateInRound_invalidNodeStateError\":\"512\",\"deadPeer\":\"64.225.69.216\",\"acceptedCBCacheMatchesAcceptedSize\":\"true\",\"awaitingForAcceptance\":\"1164\",\"snapshotHeightIntervalConditionMet\":\"5544\",\"nodeStartDate\":\"2020-08-27T04:27:52.879Z\",\"writeSnapshot_success\":\"5544\",\"snapshotWriteToDisk_success\":\"5544\",\"channelCount\":\"0\",\"observationService_unknown_size\":\"0\",\"alias\":\"Aldebaran\",\"version\":\"2.14.0\",\"addressCount\":\"924\",\"snapshotAttempt_success\":\"5544\",\"id\":\"df3f2a00d445849676e628ece489e8d8f6fc9d49a6c0f37078be802b6f3ce4a01ead6bcf0820fb66bad23ca6a2bca8c2200c9fa21b6a69680b9b2de00164d1ef\",\"transactionService_accepted_size\":\"11448\",\"checkpointsAcceptedWithDummyTxs\":\"768164\",\"downloadFinished_total\":\"1\",\"nextSnapshotHeight\":\"829738\",\"allowedForAcceptance\":\"0\",\"nodeCurrentDate\":\"2020-08-28T13:26:13.588Z\",\"rewards_selfBalanceAfterReward\":\"57818454587106\",\"snapshotCount\":\"5544\",\"observationService_inConsensus_size\":\"0\",\"TPS_last_10_seconds\":\"10.5\",\"addPeerWithRegistrationSymmetric_success\":\"1\",\"balancesBySnapshot\":\"DAG0HbSL 161899029311189, DAG0RejM 219126882700366, DAG0Ufhr 476499227376040, DAG0VBzt 47594338197598, DAG0vyN4 357019377435316, DAG1ARrj 82623024317666, DAG1KUCC 486351620292279, DAG1U4Ld 467885417079665, DAG1WwnS 476560955771088, DAG1bh6Q 468093844296592, DAG1hQvN 483126862787894, DAG1tQhf 85207888297637, DAG2Latw 79232852479506, DAG2Mt4J 84953342875593, DAG2YPmD 140865891050818, DAG2meii 51101422292767, DAG2o2Jg 408050686784717, DAG2wvkG 84906312324897, DAG3CMEY 159758533690434, DAG3ona1 60766838482254, DAG4DG1y 180573262597115, DAG4Whuj 245368011131032, DAG4kDdn 111275431517784, DAG4pmdv 77062095495741, DAG4sySb 48864692767068, DAG53v6t 244087569027288, DAG5DPxK 56821324929511, DAG5fSp3 231689263768414, DAG5sxJE 74681230372251, DAG5vU9f 48954699311812, DAG63PZk 314292902809122, DAG6Dx2X 183264242887995, DAG6EgC5 407088725752556, DAG6GRgf 41684447250928, DAG6HgDk 112548159895472, DAG6JKts 46141155421425, DAG6iPGQ 58226875080633, DAG6moTU 402886277939493, DAG6rx7T 225183561091600, DAG7yyXi 418326689973620, DAG8LX6H 45458658866517, DAG8MURx 59733503555033, DAG8d6Zo 85069643327148, DAG8uiEi 161397554387648\",\"rewards_lastRewardedHeight\":\"829708\",\"missingParents\":\"40596\",\"redownload_lastSentHeight\":\"-1\",\"deleteSnapshotInfo_success\":\"4589\",\"minTipHeight\":\"829758\",\"consensus_startOwnRound_consensusError\":\"466\",\"checkpointTipsIncremented\":\"11974\",\"heightCalculationParentMissing\":\"45\",\"consensus_startOwnRound\":\"22025\",\"address\":\"DAG5DPxKXyUvrQP4dz4HHmbutviQL1V6Y2N49A2y\",\"checkpointValidationSuccess\":\"768191\",\"redownload_maxAcceptedSnapshotHeight\":\"829736\",\"consensus_startOwnRound_unknownError\":\"2\",\"acceptedCBSinceSnapshot\":\"3675\",\"checkpointAccepted\":\"919170\",\"acceptMajorityCheckpointBlockSelectedCount_3\":\"45830\",\"transactionAcceptedFromRedownload\":\"452940\",\"badPeerAdditionAttempt\":\"50\",\"nodeCurrentTimeMS\":\"1598621173588\",\"batchTransactionsEndpoint\":\"838\",\"consensus_startOwnRound_consensusStartError\":\"8860\",\"rewards_selfSnapshotReward\":\"673400673\",\"metricsRound\":\"11871\",\"consensus_startOwnRound_noTipsForConsensusError\":\"6\",\"genesisHash\":\"e2c7316729b426ea15a4471b9d742ef42e57b68a85606e597d0a40385203e834\",\"peers\":\"4d7ce API: http://142.93.177.241:9000 --- 710b3 API: http://54.177.239.156:9000 --- 0619f API: http://167.99.132.22:9000 --- 130f1 API: http://35.245.174.85:9000 --- 754a4 API: http://95.216.7.82:9000 --- aa336 API: http://167.99.25.87:9000 --- 278ac API: http://128.199.72.232:9000 --- 3b2a1 API: http://149.28.206.238:9000 --- d0248 API: http://128.199.178.90:9000 --- e5272 API: http://45.32.218.171:9000 --- 028a3 API: http://64.225.69.216:9000 --- 3fe99 API: http://54.185.204.197:9000 --- c9078 API: http://64.227.110.71:9000 --- f6baa API: http://198.211.96.168:9000 --- 7bbda API: http://165.22.70.165:9000 --- 3a39e API: http://167.172.4.196:9000 --- c9c4e API: http://161.35.97.20:9000 --- 70e08 API: http://134.209.195.37:9000 --- c4e8c API: http://64.225.46.57:9000 --- bdb18 API: http://144.202.37.59:9000 --- ea4f1 API: http://157.245.173.182:9000 --- 62988 API: http://18.144.132.60:9000 --- effb0 API: http://161.35.226.110:9000 --- b8f7d API: http://45.63.14.140:9000 --- f0073 API: http://178.128.130.251:9000 --- 952d6 API: http://52.13.130.169:9000 --- 87c8b API: http://34.91.142.112:9000 --- 4c1ba API: http://34.87.248.47:9000 --- e0c1e API: http://54.219.13.163:9000 --- 2afe7 API: http://34.91.201.98:9000 --- 4c015 API: http://35.237.233.84:9000 --- 97e44 API: http://142.93.227.100:9000 --- e0c15 API: http://167.172.228.202:9000 --- eb2b7 API: http://68.183.203.17:9000 --- fb399 API: http://45.77.3.10:9000 --- f3d72 API: http://184.169.207.196:9000 --- 4bc6c API: http://45.77.65.16:9000 --- 3e16c API: http://35.243.183.166:9000 --- aa172 API: http://34.73.82.172:9000 --- c41ca API: http://167.172.111.46:9000 --- 2ed28 API: http://44.230.198.81:9000 --- 708b4 API: http://174.138.47.80:9000 --- 58f11 API: http://167.172.4.203:9000\",\"rewards_snapshotCount\":\"5126\",\"reDownloadFinished_total\":\"39\",\"syncBufferSize\":\"0\",\"lastSnapshotHeight\":\"829736\",\"observationService_accepted_size\":\"0\",\"observationService_pending_size\":\"0\",\"heightBelow\":\"2\",\"transactionService_pending_size\":\"0\",\"rewards_stardustSnapshotReward\":\"3292181069\",\"deleteSnapshot_success\":\"4589\",\"snapshotHeightIntervalConditionNotMet\":\"5708\",\"checkpointAcceptBlockAlreadyStored\":\"1412716\",\"activeTips\":\"19\",\"transactionService_inConsensus_size\":\"12\"}}")

	d := MetricsEnvelope{}

	e := json.Unmarshal(jsonBytes, &d)

	if e != nil {
		t.Error(e)
	}
}
