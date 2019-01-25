base.img:
	wget -O $@

image.img: base.img
	docker run --rm --privileged -v $(shell pwd):/app ubuntu /app/customize.sh $(SITE_NAME) $(SITE_PORT)
