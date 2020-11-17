package signer

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"

	"github.com/fabricorgi/cmd/orgchecker"
)

//Перехватчик ошибок из CMD
func errorHandler(cmd *exec.Cmd) error {
	cmd.Stderr = os.Stderr
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

//SignAndAdd запуск ряда комманд для получения конфигурационного блока и внесения изменений
func SignAndAdd(payload *orgchecker.OrganizationConfig, channel string) error {

	// Инициализация переменной cmd
	var cmd *exec.Cmd
	var err error

	// Цикл по добавлению организации в канал
	cmd = exec.Command("bash", "-c", fmt.Sprintf("peer channel fetch config config_block.pb --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -o ${FABRICORGI_ORDERER_IP} -c %s", channel))
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", fmt.Sprintf(`jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"'%s'":.[1]}}}}}' config.json ./organisation.json > modified_config.json`, payload.Policies.Readers.Policy.Value.Identities[0].Principal.MspIdentifier))
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input config.json --type common.Config --output config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", fmt.Sprintf("configtxlator compute_update --channel_id %s --original config.pb --updated modified_config.pb --output organisation_update.pb", channel))
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input organisation_update.pb --type common.ConfigUpdate | jq . > organisation_update.json")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", fmt.Sprintf(`echo '{"payload":{"header":{"channel_header":{"channel_id":"%s", "type":2}},"data":{"config_update":'$(cat organisation_update.json)'}}}' | jq . > organisation_update_in_envelope.json`, channel))
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input organisation_update_in_envelope.json --type common.Envelope --output organisation_update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "peer channel signconfigtx --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f organisation_update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", fmt.Sprintf("peer channel update --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f organisation_update_in_envelope.pb -c %s -o ${FABRICORGI_ORDERER_IP}", channel))
	err = errorHandler(cmd)

	if err != nil {
		return err
	}

	return nil
}

// SignAndRemove функция для удаления организации из канала
func SignAndRemove(payload *orgchecker.OrganizationRemove) error {

	// Инициализация переменной cmd
	var cmd *exec.Cmd
	var err error

	// Цикл по удалению организации из канала
	cmd = exec.Command("bash", "-c", "peer channel fetch config config_block.pb --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -o ${FABRICORGI_ORDERER_IP} -c mainchannel")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", fmt.Sprintf("jq '.channel_group.groups.Application.groups.%s)' config.json > modified_config.json", payload.OrgName))
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input config.json --type common.Config --output config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator compute_update --channel_id mainchannel --original config.pb --updated modified_config.pb --output update.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input update.pb --type common.ConfigUpdate | jq . > update.json")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", `echo '{"payload":{"header":{"channel_header":{"channel_id":"mainchannel", "type":2}},"data":{"config_update":'$(cat update.json)'}}}' | jq . > update_in_envelope.json`)
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input update_in_envelope.json --type common.Envelope --output update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "peer channel signconfigtx --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "peer channel update --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f update_in_envelope.pb -c mainchannel -o ${FABRICORGI_ORDERER_IP}")
	err = errorHandler(cmd)

	if err != nil {
		return err
	}

	return nil
}

func SignAndChangeConfig(payload *orgchecker.OrdererConfig) error {

	// Инициализация переменной cmd
	var cmd *exec.Cmd
	var err error

	// Цикл по удалению организации из канала
	cmd = exec.Command("bash", "-c", "peer channel fetch config config_block.pb --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -o ${FABRICORGI_ORDERER_IP} -c mainchannel")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json")
	err = errorHandler(cmd)

	if payload.BatchSizeMaxMessageCount != 0 {
		cmd = exec.Command("bash", "-c", fmt.Sprintf("jq '.channel_group.groups.Orderer.values.BatchSize.value.max_message_count = %d' config.json > modified_config.json", payload.BatchSizeMaxMessageCount))
		err = errorHandler(cmd)
	}

	if payload.BatchSizeAbsoluteMaxBytes != 0 {
		cmd = exec.Command("bash", "-c", fmt.Sprintf("jq '.channel_group.groups.Orderer.values.BatchSize.value.absolute_max_bytes = %d' config.json > modified_config.json", payload.BatchSizeAbsoluteMaxBytes*1024*1024))
		err = errorHandler(cmd)
	}

	if payload.BatchSizePrefferedMaxBytes != 0 {
		cmd = exec.Command("bash", "-c", fmt.Sprintf("jq '.channel_group.groups.Orderer.values.BatchSize.value.preferred_max_bytes = %v' config.json > modified_config.json", math.Round(payload.BatchSizePrefferedMaxBytes*1024*1024)))
		err = errorHandler(cmd)
	}

	if payload.BatchTimeout != "" {
		cmd = exec.Command("bash", "-c", fmt.Sprintf(`jq '.channel_group.groups.Orderer.values.BatchTimeout.value.timeout = "%s"' config.json > modified_config.json`, payload.BatchTimeout))
		err = errorHandler(cmd)
	}
	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input config.json --type common.Config --output config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator compute_update --channel_id mainchannel --original config.pb --updated modified_config.pb --output update.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input update.pb --type common.ConfigUpdate | jq . > update.json")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", `echo '{"payload":{"header":{"channel_header":{"channel_id":"mainchannel", "type":2}},"data":{"config_update":'$(cat update.json)'}}}' | jq . > update_in_envelope.json`)
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input update_in_envelope.json --type common.Envelope --output update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "peer channel signconfigtx --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "CORE_PEER_LOCALMSPID="+fmt.Sprintf("%sMSP", os.Getenv("ORG_NAME"))+";CORE_PEER_MSPCONFIGPATH="+fmt.Sprintf("/shared/admin%sMSP", strings.ToTitle(os.Getenv("ORG_NAME")))+"; peer channel update --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f update_in_envelope.pb -c mainchannel -o ${FABRICORGI_ORDERER_IP}")
	err = errorHandler(cmd)

	if err != nil {
		return err
	}

	return nil
}
